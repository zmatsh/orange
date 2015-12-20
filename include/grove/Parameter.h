/*
** Copyright 2014-2015 Robert Fratto. See the LICENSE.txt file at the top-level
** directory of this distribution.
**
** Licensed under the MIT license <http://opensource.org/licenses/MIT>. This file
** may not be copied, modified, or distributed except according to those terms.
*/

#pragma once

#include "Expression.h"
#include "Named.h"
#include "Accessible.h"

class Type;

class Parameter : public Expression, public Named, public Accessible {
public:
	virtual llvm::Value* getPointer() const override;
	
	virtual bool hasPointer() const override;
	
	virtual llvm::Value* getValue() const override;
	
	virtual ASTNode* copy() const override;
	
	virtual bool isAccessible() const override;
	
	virtual Expression* access(OString name, const ASTNode* hint)
		const override;
	
	Parameter(Type* type, OString name);
};
