/*
** Copyright 2014-2015 Robert Fratto. See the LICENSE.txt file at the top-level
** directory of this distribution.
**
** Licensed under the MIT license <http://opensource.org/licenses/MIT>. This file
** may not be copied, modified, or distributed except according to those terms.
*/

#pragma once 

#include "Block.h"
#include "Named.h"
#include "Typed.h"

/**
 * Represents a class declaration.
 */
class ClassDecl : public Block, public Named, public Typed
{
	
};