#include "gen/ReturnExpr.h"
#include "gen/generator.h"
#include "gen/VarExpr.h"

void ReturnExpr::resolve() {
	if (resolved)
		return;

	if (expr) {
		expr->resolve();
	}

	resolved = true;
}

Value* ReturnExpr::Codegen() {
	BasicBlock *bb = CG::Symtab->getFunctionEnd();
	if (bb == nullptr) {
		std::cerr << "fatal: no Function End found!\n";
		exit(1);
	}

	if (expr && CG::Symtab->getRetVal() == nullptr) {
		std::cerr << "fatal: no return value found (symtab " << CG::Symtab->ID << ") !\n"; 
		exit(1);
	}

	if (expr) {
		DEBUG_MSG("ReturnExpr: STARTING CODEGEN FOR EXPR");
		Value *v = expr->Codegen();

		if (returnsPtr(expr->getClass())) {
			v = CG::Builder.CreateLoad(v);
		}
		
		CG::Builder.CreateStore(v, CG::Symtab->getRetVal());
		Value *r = CG::Builder.CreateBr(bb);
		return r; 
	} else {
		Value *r = CG::Builder.CreateBr(bb);
		return r;
	}
}
