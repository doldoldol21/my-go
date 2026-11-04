const vscode = require('vscode');

// vscode://doldoldol21.open-url?url=<인코딩된 주소> → 내장 브라우저로 연다
exports.activate = (context) => {
	context.subscriptions.push(vscode.window.registerUriHandler({
		handleUri(uri) {
			const url = new URLSearchParams(uri.query).get('url');
			if (url) vscode.commands.executeCommand('workbench.action.browser.open', url);
		},
	}));
};
