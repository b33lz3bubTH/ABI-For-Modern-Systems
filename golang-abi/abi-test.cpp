#include <iostream>
#include <dlfcn.h>  // For dlopen, dlsym, dlclose
#include <cstdlib>  // For free

// Define a function pointer type that returns a double and takes a const char* argument
typedef char* (*AbiFunction)(const char*);

int main() {
    // Path to the compiled .so file
    const char* lib_path = "/tmp/lib_crypto.so";

    // Load the shared library
    void* handle = dlopen(lib_path, RTLD_LAZY);
    if (!handle) {
        std::cerr << "Failed to load library: " << dlerror() << std::endl;
        return 1;
    }

    // Load the ABI function
    AbiFunction encrypt = (AbiFunction)dlsym(handle, "encrypt");
    if (!encrypt) {
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
    char* result = encrypt(json_str);

    // Output the result
    std::cout << "Computed result: " << result << std::endl;



    const char* data_json_str = R"({
    "data_encrypted": "JxpXsMv75Wxrp+rTdYh+rDZ/Pyr7/je8ydlooHktgi2HNWqUCdoAYbwb9oVuhY1y/nu1bzCwgQ+V0tMnFFThbtkwUHBjDc9Ma/4wRWWRCT/JGLl+GMCigX/VwO43uWwcP67A/x64CiNVvGbunDxrGhK+ccHEg8S7pXt8k0keoooEkz5Ba9o8uaO6Jjg2Hui7ZPO3noroREni4TlaaAntdeIvrHvd27sLPqnEPk7zbj19ibPE5iEe9QiGkaoUxJBpDxd4Z44gV2UQZWaj7dmrc9CB1OrZ8svhKGjNJ4js/nPOrU5XWNEWyZ7HzHZFnCUPBxNaFFqm3Q8UR0wD9U0Xw9bAYmTRcQ9c+SFO3gMCQ2nWGHqpfu8ojcurwQa7KMBjPVPGIPB+ndyOO5KEDpHGfKIIPsvmEaUQclbEyMYtU5eP5N8fBWyGl92W7avYOZWkSdmzKmTbIWS5SAmGNtwiQIBCQ3RuB0oseZFFCq9WXwGL3u1vcxV1Dcjd/IEK/rKcmQiu8NNPEZ1rMIBCuHTI5w9aOkOCjyYW1iMM84TyW7Shz9IJTcI7lUnA434fvhIm+CJxMV2P1n8a9r/+4f4wd5IRURmVWwj/HKOA2+m5BJsCm/FX1spL4bdXr2GKlQg2x9MP9JJL+ffKxw==",
    "aes_key": "k+fWGto3KWYSkemO57b5YuY0w1I3pOTEzvsbqNqw1Kk=",
    "sha512": "EcgCb8egOOCfzXT6Ktx5jpWu6Ef92QkdlT1PW1UmLFh+iROYa921IJT0V9G5pzKy7WmB5QQcsqnZ3K/yoSbTxg=="
    })";

    std::cout << "Decrypting Data" << data_json_str << std::endl;


    AbiFunction decrypt = (AbiFunction)dlsym(handle, "decrypt");
    if (!decrypt) {
        std::cerr << "Failed to find function: " << dlerror() << std::endl;
        dlclose(handle);
        return 1;
    }


    char* data_response = decrypt(data_json_str);

    // Output the result
    std::cout << "Decrypted result: " << data_response << std::endl;

    // Close the library
    dlclose(handle);

    return 0;
}


// abi(master) : g++ -o abi-test.out abi-test.cpp -ldl
