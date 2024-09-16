# Abi Build with rust

========================

1. cargo build --release
2. g++ -o abi-test.out abi-test.cpp -ldl
3. ./abi-test.out

# Abi Build with Golang

=====================================

1. go build -buildmode=c-shared -o /tmp/lib_crypto.so
2. chmod +r /tmp/lib_crypto.so
3. g++ -o abi-test.out abi-test.cpp -ldl
4. ./abi-test.out

## output

`{"data_encrypted":"bH7gGnWJ3hn6ttqcQMEQlCsBswX67h6rkdRbVTmaay0Sn7JglRkiqoJVi0gRkjgnn+nfhMUELFxb6G8GG0xP5V6tDpw1yRhegtRWKxjBwrML+D88B36CR3aT4VFDNujA5naCbb+5qd0kFYK8X1SM8x4D0jdm/wkBTaUAIpiMxECUWzlpaE84fguwE+0N1IrQtmddt51Vn6dtXV+++GGmHVwFjCR7d4SdpuaJkZVRBwY18gj5b7ZZ+XrafKZNfcNdMemPa9d+EjDhgWTYLu/BxDAMKZZNAGbtA0vM6pE2y4xzdNZd9Fj/+kVb939ekjt0eGwEDlvTVW8L6hFlBTKjD65zSMKHmj3BPLh4cijds3smBWv1puneyLn1gZBnkIStzSi+ufTGLKzp02E6wZGhXdnnYl6uBjcB5mYipH9dHagMRKjiCgd8lBHSE3zDnzdPzQUzOES+F81Q01CnZzEtvlYt+33Aq5XYODNwlZkZFOczsi04jcA02qKxR7V5bwwStoOYNyN5wHNf0aXGTKle5ZMfhvvxQAMssTw9of3C1Yr6838In8f8yPLmyPl8ldHriiZBA+JOd7sIJDrs6pwbG7aRJ2uHfMbYogyoI1Ye+/iIZtyc7XxqQmHFUmxCs38JJ4gkzPyaZSmnOg==","aes_key":"1j1OHYBnPJh/JEIpB3iGSqrsA6mJ9ruEaTIitrNdsSc=","sha512":"EcgCb8egOOCfzXT6Ktx5jpWu6Ef92QkdlT1PW1UmLFh+iROYa921IJT0V9G5pzKy7WmB5QQcsqnZ3K/yoSbTxg=="}`
