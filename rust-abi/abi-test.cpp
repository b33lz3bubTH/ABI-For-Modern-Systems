#include <iostream>
#include <dlfcn.h>  // For dlopen, dlsym, dlclose
#include <cstdlib>  // For free

// Define a function pointer type that returns a double and takes a const char* argument
typedef double (*AbiFunction)(const char*);

int main() {
    // Path to the compiled .so file
    const char* lib_path = "./target/release/libpricing_strategy.so";

    // Load the shared library
    void* handle = dlopen(lib_path, RTLD_LAZY);
    if (!handle) {
        std::cerr << "Failed to load library: " << dlerror() << std::endl;
        return 1;
    }

    // Load the ABI function
    AbiFunction compute_price_statistics = (AbiFunction)dlsym(handle, "compute_price_statistics");
    if (!compute_price_statistics) {
        std::cerr << "Failed to find function: " << dlerror() << std::endl;
        dlclose(handle);
        return 1;
    }

    // Prepare input data as a JSON string
    const char* json_str = R"({
        "prices": [
            { "price": 500, "name": "pandu ranga" },
            { "price": 500 },
            { "price": 450 },
            { "price": 450 },
            { "price": 470 },
            { "price": 470 },
            { "price": 600 },
            { "price": 700 },
            { "price": 480 },
            { "price": 480 },
            { "price": 480 },
            { "price": 650, "name": "Product x" }
        ],
        "my_product_price": 3000
    })";

    // Call the ABI function and store the result
    double result = compute_price_statistics(json_str);

    // Output the result
    std::cout << "Computed result: " << result << std::endl;

    // Close the library
    dlclose(handle);

    return 0;
}


// abi(master) : g++ -o abi-test.out abi-test.cpp -ldl
