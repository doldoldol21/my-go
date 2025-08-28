package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/doldoldol21/my-go/project/todo/snowflake"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// 할일 목록 제한
const MAX_TODO_SIZE = 100
const MAX_CONTENT_LETTER_SIZE = 50
const MAX_PRIORITY = 999

var (
	content        string
	targetID       int64
	targetPriority int16
)

// Cmd
var rootCmd = &cobra.Command{
	Short: "A simple TODO CLI application",
	Long:  "A simple command line TODO application built with Go and Cobra",
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new TODO",
	Run:   add,
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete TODO",
	Run:   delete,
}

var updatePriorityCmd = &cobra.Command{
	Use:   "update-priority",
	Short: "Update priority of TODO",
	Run:   updatePriority,
}

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Stamping done mark to TODO!",
	Run:   done,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View all TODO list",
	Run:   list,
}

type Todo struct {
	ID           int64
	Content      string
	Priority     int16
	CreationTime time.Time
	UpdateTime   time.Time
	IsDone       bool
}

type TodoList struct {
	Todos []Todo
}

func (tl *TodoList) addTodo(todo Todo) {
	tl.Todos = append(tl.Todos, todo)
}

func (tl *TodoList) isFullTodo() bool {
	return len(tl.Todos) == MAX_TODO_SIZE
}

// 우선순위 기준 정렬 (낮은 숫자가 높은 우선순위)
func list(cmd *cobra.Command, args []string) {
	todoList := getTodoList()
	sort.Slice(todoList.Todos, func(i, j int) bool {
		pre := todoList.Todos[i]
		nex := todoList.Todos[j]
		if pre.Priority != nex.Priority {
			return pre.Priority < nex.Priority
		}
		return pre.CreationTime.Before(nex.CreationTime)
	})
	format := "2006-01-02 15:04:05"
	data := [][]string{
		{"ID", "Content", "Priority", "CreationTime", "UpdateTime", "Done"},
	}
	doneStatus := map[bool]string{true: "✅", false: "❌"}
	for _, todo := range todoList.Todos {
		row := []string{
			fmt.Sprintf("%d", todo.ID),
			todo.Content,
			fmt.Sprintf("%d", todo.Priority),
			todo.CreationTime.Local().Format(format),
			todo.UpdateTime.Local().Format(format),
			doneStatus[todo.IsDone],
		}
		data = append(data, row)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data[0])
	table.Bulk(data[1:])
	table.Render()
}

func add(cmd *cobra.Command, args []string) {
	contentLen := len(content)
	if contentLen > MAX_CONTENT_LETTER_SIZE {
		fmt.Printf("maximum letter count is %d. input letter count is %d\n", MAX_CONTENT_LETTER_SIZE, contentLen)
		return
	}
	todoList := getTodoList()
	if isFullTodo := todoList.isFullTodo(); isFullTodo {
		fmt.Fprintln(os.Stdout, "Todo List is Full!")
		return
	}
	todoID, err := snowflake.NextId()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
	// 자동 우선순위 할당 (기존 할일 개수 + 1)
	priority := int16(len(todoList.Todos) + 1)
	if priority > MAX_PRIORITY {
		priority = MAX_PRIORITY
	}
	now := time.Now().UTC()
	todoList.addTodo(Todo{
		ID:           todoID,
		Content:      content,
		Priority:     priority,
		CreationTime: now,
		UpdateTime:   now,
		IsDone:       false,
	})
	writeTodoFile(todoList)
}

func delete(cmd *cobra.Command, args []string) {
	todoList := getTodoList()
	for i, todo := range todoList.Todos {
		if todo.ID == targetID {
			todoList.Todos = append(todoList.Todos[:i], todoList.Todos[i+1:]...)
			writeTodoFile(todoList)
			return
		}
	}
	fmt.Fprintf(os.Stderr, "error: TODO not exists")
}

func updatePriority(cmd *cobra.Command, args []string) {
	if targetPriority > MAX_PRIORITY {
		fmt.Fprintf(os.Stderr, "error: priority can't bigger then %d\n", MAX_PRIORITY)
		return
	}
	todoList := getTodoList()
	for i, todo := range todoList.Todos {
		if todo.ID == targetID {
			todoList.Todos[i].Priority = targetPriority
			todoList.Todos[i].UpdateTime = time.Now().UTC()
			writeTodoFile(todoList)
			return
		}
	}
	fmt.Fprintf(os.Stderr, "error: TODO not fount\n")
}

// 완료 상태 토글 (--unmark 플래그로 완료 해제)
func done(cmd *cobra.Command, args []string) {
	unmarkFlag := cmd.Flags().Lookup("unmark")
	isDone := !unmarkFlag.Changed
	todoList := getTodoList()
	for i, todo := range todoList.Todos {
		if todo.ID == targetID {
			todoList.Todos[i].IsDone = isDone
			todoList.Todos[i].UpdateTime = time.Now().UTC()
			writeTodoFile(todoList)
			return
		}
	}
	fmt.Fprintf(os.Stderr, "error: TODO not fount\n")
}

func writeTodoFile(todoList *TodoList) {
	todoFilePath := getTodoFilePath()
	b, err := json.Marshal(todoList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(todoFilePath, b, 0666); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Done !")
}

func getTodoList() *TodoList {
	var todoList TodoList
	todoFilePath := getTodoFilePath()
	todoFile, err := os.ReadFile(todoFilePath)
	if err != nil {
		return &todoList
	}
	if len(todoFile) > 0 {
		if err := json.Unmarshal(todoFile, &todoList); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n (you can delete todo.json)\n", err)
			os.Exit(1)
		}
	}
	return &todoList
}

func getTodoFilePath() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return filepath.Join(wd, "todo.json")
}

func init() {
	addCmd.Flags().StringVarP(&content, "content", "c", "", fmt.Sprintf("TODO content (required)\n maximum letter count is %d", MAX_CONTENT_LETTER_SIZE))
	addCmd.MarkFlagRequired("content")

	deleteCmd.Flags().Int64VarP(&targetID, "id", "i", 0, "TODO ID (required)")
	deleteCmd.MarkFlagRequired("id")

	updatePriorityCmd.Flags().Int64VarP(&targetID, "id", "i", 0, "TODO ID (required)")
	updatePriorityCmd.Flags().Int16VarP(&targetPriority, "priority", "p", 0, "Update Priority number to (required)")
	updatePriorityCmd.MarkFlagRequired("id")
	updatePriorityCmd.MarkFlagRequired("priority")

	doneCmd.Flags().Int64VarP(&targetID, "id", "i", 0, "TODO ID (required)")
	doneCmd.Flags().BoolP("unmark", "u", false, "Unmark TODO")
	doneCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(addCmd, deleteCmd, updatePriorityCmd, doneCmd, listCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		// fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(0)
	}
}
