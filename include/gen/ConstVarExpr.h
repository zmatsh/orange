#ifndef __CONST_VAR_EXPR_H__
#define __CONST_VAR_EXPR_H__
#include "AST.h"

class ConstVarExpr : public Expression {
public:
	virtual std::string getClass() { return "ConstVarExpr"; }
	
	AnyType *type; 
	std::string name;
	Expression *value;

	Value* Codegen();

	virtual bool isSigned() { return type->isSigned(); }

	virtual Type* getType() {
		return type->getType();
	}

	virtual std::string string() {
		return name;
	}

	virtual Statement* clone() { 
		std::string* name_p = new std::string(name); 
		ConstVarExpr* ret = new ConstVarExpr((AnyType *)type->clone(), name_p, (Expression*)value->clone());
		delete name_p; 
		return ret;
	}

	virtual bool returnsPtr() { return true; }

	virtual bool isConstant() { return false; }

	virtual void resolve();

	ConstVarExpr(AnyType *type, std::string *name, Expression *value);

	~ConstVarExpr();
};

#endif 