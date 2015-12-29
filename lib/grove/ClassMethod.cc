/*
** Copyright 2014-2015 Robert Fratto. See the LICENSE.txt file at the top-level
** directory of this distribution.
**
** Licensed under the MIT license <http://opensource.org/licenses/MIT>. This file
** may not be copied, modified, or distributed except according to those terms.
*/

#include <grove/ClassMethod.h>
#include <grove/ClassDecl.h>
#include <grove/Parameter.h>

#include <grove/types/Type.h>
#include <grove/types/ReferenceType.h>

#include <util/assertions.h>
#include <util/copy.h>

ProtectionLevel ClassMethod::defaultProtectionLevel() const
{
	return ProtectionLevel::PROTECTION_PUBLIC;
}

void ClassMethod::findDependencies()
{
	Function::findDependencies();
	
	addDependency(m_class);
}

ASTNode* ClassMethod::copy() const
{
	return new ClassMethod(*this);
}

std::vector<ObjectBase**> ClassMethod::getMemberNodes()
{
	auto list = Function::getMemberNodes();
	list.insert(list.end(), {
		(ObjectBase **)&m_class,
		(ObjectBase **)&m_this_param
	});
	return list;
}

std::vector<std::vector<ObjectBase *>*> ClassMethod::getMemberLists()
{
	return Function::getMemberLists();
}

Parameter* ClassMethod::getThisParam() const
{
	return m_this_param;
}

ClassMethod::ClassMethod(OString name, ClassDecl* theClass,
						 std::vector<Parameter *> params)
: Function(name, params)
{
	assertExists(theClass, "ClassMethod created with no class");
	
	auto ty = new ReferenceType(theClass);
	m_this_param = new Parameter(ty->getPointerTo(), "this");
	if (m_params.size() == 0)
	{
		addChild(m_this_param);
	}
	else
	{
		addChild(m_this_param, m_params.at(0), 0);
	}
	
	m_params.insert(m_params.begin(), m_this_param);
	
	m_class = theClass;
}

ClassMethod::ClassMethod(const ClassMethod& other)
: Function(other.m_name, copyVector(other.getParams()))
{
	m_this_param = getParams().at(0);
	m_class = other.m_class;
	
	other.defineCopy(this);
}