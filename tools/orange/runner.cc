#include <orange/runner.h>
#include <orange/generator.h>

Runner::Runner(std::string pathname) {
	m_pathname = pathname;

	bool added = GeneratingEngine::sharedEngine()->addRunner(this);
	if (added == false) {
		throw std::runtime_error("File cannot be added as an entity twice.");
	}
}

void Runner::haltRun() {
	m_isRunning = false;
}

bool Runner::hasError() {
	for (auto msg : m_messages) {
		if (msg.type() == ERROR) return true; 
	}

	return false;
}

RunResult Runner::run() {
	// First, set us to running and activate us as the current runner.
	m_isRunning = true;
	GeneratingEngine::sharedEngine()->setActive(this);

	// Try to find our file...
	FILE *file = fopen(pathname().c_str(), "r");
	if (file == nullptr) {
		// Halt our current run and return an error.
		haltRun();

		CompilerMessage msg(NO_FILE, "file " + pathname() + " not found.", pathname(), -1, -1, -1, -1);
		return RunResult(pathname(), false, 1, msg);
	}

	// Parse the file. get yyin and yyparse and use them
	extern FILE* yyin;
	extern int yyparse();

	yyin = file;
	yyparse();

	// TODO: run the code.
	int retCode = 0;

	// Do cleanup.
	fclose(file);
	m_isRunning = false;

	bool succeeded = (retCode == 0) && (hasError() == false); 
	return RunResult(pathname(), succeeded, retCode, m_messages);
}

void Runner::log(CompilerMessage message) {
	// We only want to log if we're currently running.
	if (m_isRunning == false) return;

	m_messages.push_back(message);
}

std::string Runner::pathname() const {
	return m_pathname; 
}