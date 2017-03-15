//============================================================================
// Name        : DPU_Lua_Translator.cpp
// Author      : test
// Version     :
// Copyright   : Your copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <lua.hpp>
#include <iostream>
#include <string>

using namespace std;

int main() {
	cout << "LUA reading test" << endl << endl;
	lua_State *L = luaL_newstate(); //lua_open();

	//luaL_openlibs(L); //FAIL no console output?
	// load Lua libraries
	static const luaL_Reg lualibs[] =
			{ { "base", luaopen_base }, { NULL, NULL } };

	const luaL_Reg *lib = lualibs;
	for (; lib->func != NULL; lib++) {
		lib->func(L);
		lua_settop(L, 0);
	}


	luaL_dostring(L, "print(\"Lua string test\\nNext line test\\n\")");

	if (!luaL_dofile(L, "script.lua")) {
		lua_getglobal(L, "x");
		int x = lua_tointeger(L, -1);
		cout << "Lua Variable test: " << x << endl;
	}

	lua_close(L);
	cout << "Press any key!";
	(const char *) cin.get();

	return 0;
}
