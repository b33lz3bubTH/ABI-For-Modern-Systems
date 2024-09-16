import ctypes
import json

# Define the path to your shared library
lib_path = '/tmp/lib_crypto.so'

# Load the shared library
try:
    lib = ctypes.CDLL(lib_path)
except OSError as e:
    print(f"Failed to load library: {e}")
    exit(1)

# Define the function signatures for the ABI functions
lib.encrypt.argtypes = [ctypes.c_char_p]
lib.encrypt.restype = ctypes.c_char_p

lib.decrypt.argtypes = [ctypes.c_char_p]
lib.decrypt.restype = ctypes.c_char_p

# Prepare input data as a JSON string
input_data = {
    "prices": [
        { "price": 500, "name": "Ranga 1000000" },
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
}

# Convert the input data to a JSON string
json_str = json.dumps(input_data)

# Call the encrypt function
encrypted_result = lib.encrypt(json_str.encode('utf-8'))
print("Encrypted result:", ctypes.c_char_p(encrypted_result).value.decode('utf-8'))

# Prepare encrypted data as a JSON string (sample data for decryption)
encrypted_data = {
    "data_encrypted": "JxpXsMv75Wxrp+rTdYh+rDZ/Pyr7/je8ydlooHktgi2HNWqUCdoAYbwb9oVuhY1y/nu1bzCwgQ+V0tMnFFThbtkwUHBjDc9Ma/4wRWWRCT/JGLl+GMCigX/VwO43uWwcP67A/x64CiNVvGbunDxrGhK+ccHEg8S7pXt8k0keoooEkz5Ba9o8uaO6Jjg2Hui7ZPO3noroREni4TlaaAntdeIvrHvd27sLPqnEPk7zbj19ibPE5iEe9QiGkaoUxJBpDxd4Z44gV2UQZWaj7dmrc9CB1OrZ8svhKGjNJ4js/nPOrU5XWNEWyZ7HzHZFnCUPBxNaFFqm3Q8UR0wD9U0Xw9bAYmTRcQ9c+SFO3gMCQ2nWGHqpfu8ojcurwQa7KMBjPVPGIPB+ndyOO5KEDpHGfKIIPsvmEaUQclbEyMYtU5eP5N8fBWyGl92W7avYOZWkSdmzKmTbIWS5SAmGNtwiQIBCQ3RuB0oseZFFCq9WXwGL3u1vcxV1Dcjd/IEK/rKcmQiu8NNPEZ1rMIBCuHTI5w9aOkOCjyYW1iMM84TyW7Shz9IJTcI7lUnA434fvhIm+CJxMV2P1n8a9r/+4f4wd5IRURmVWwj/HKOA2+m5BJsCm/FX1spL4bdXr2GKlQg2x9MP9JJL+ffKxw==",
    "aes_key": "k+fWGto3KWYSkemO57b5YuY0w1I3pOTEzvsbqNqw1Kk=",
    "sha512": "EcgCb8egOOCfzXT6Ktx5jpWu6Ef92QkdlT1PW1UmLFh+iROYa921IJT0V9G5pzKy7WmB5QQcsqnZ3K/yoSbTxg=="
}

# Convert the encrypted data to a JSON string
encrypted_json_str = json.dumps(encrypted_data)

# Call the decrypt function
decrypted_result = lib.decrypt(encrypted_json_str.encode('utf-8'))
print("Decrypted result:", ctypes.c_char_p(decrypted_result).value.decode('utf-8'))

decrypted_data = json.loads(ctypes.c_char_p(decrypted_result).value.decode('utf-8'))
original_data = decrypted_data["original_data"]
print("Decrypted Formatted: ", original_data)

# Clean up
del encrypted_result
del decrypted_result
