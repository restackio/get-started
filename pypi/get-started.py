import platform
import subprocess

def detect_architecture_and_invoke():
    architecture = platform.machine()
    print(f"Detected architecture: {architecture}")

    if architecture == 'x86_64':
        # Invoke the binary for x86_64 architecture
        subprocess.run(['./binary_x86_64'])
    elif architecture == 'arm64' or architecture == 'aarch64':
        # Invoke the binary for ARM architecture
        subprocess.run(['./binary_arm64'])
    else:
        print("Unsupported architecture")

if __name__ == "__main__":
    detect_architecture_and_invoke()