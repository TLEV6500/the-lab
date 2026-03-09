cmake_minimum_required(VERSION 3.20)
project(LabCore VERSION 0.1.0 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Check if compiling for WebAssembly via Emscripten
if(EMSCRIPTEN)
    message(STATUS "Configuring for WebAssembly")
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -s USE_WEBGPU=1 -s WASM=1")
else()
    message(STATUS "Configuring for Native Desktop")
endif()

add_executable(LabCore src/main.cpp)