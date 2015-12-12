/*
** Copyright 2014-2015 Robert Fratto. See the LICENSE.txt file at the top-level
** directory of this distribution.
**
** Licensed under the MIT license <http://opensource.org/licenses/MIT>. This file
** may not be copied, modified, or distributed except according to those terms.
*/

#include <grove/MemberVarDecl.h>
#include <grove/ClassDecl.h>

#include <util/assertions.h>

unsigned int MemberVarDecl::getOffset() const
{
	auto parentClass = findParent<ClassDecl *>();
	assertExists(parentClass, "couldn't find a parent class for "
				 "MemberVarDecl");
	
	auto&& members = parentClass->getMembers();
	auto it = std::find(members.begin(), members.end(), this);

	if (it == members.end())
	{
		throw fatal_error("couldn't find member's offset in parent");
	}
	
	return (unsigned int)std::distance(members.begin(), it);
}

MemberVarDecl::MemberVarDecl(Type* type, OString name, Expression* expression)
: VarDecl(type, name, expression)
{
	// Do nothing.
}