import os
import urllib.request
import subprocess
import importlib.metadata

def _run(bash_script):
    try:
        return subprocess.call(bash_script, shell=True)
    except Exception as e:
        print(f"Error executing script: {e}")
        return 1

def restack_get_started():
    # Get system architecture and platform
    arch = os.uname().machine
    platform_name = os.uname().sysname
    version = importlib.metadata.version("restackio.get-started")
    print(f"Version: {version}")

    # Determine binary path based on platform and architecture
    if platform_name == "Darwin":
        binary_name = "restack-get-started-darwin-amd64"
        binary_url = f"https://github.com/restackio/get-started/releases/download/v{version}/{binary_name}"
    elif platform_name == "Linux" and arch == "x86_64":
        binary_name = "restack-get-started-linux-amd64"
        binary_url = f"https://github.com/restackio/get-started/releases/download/v{version}/{binary_name}"
    else:
        print("Unsupported platform")
        return 1
    
    # Download binary
    try:
        urllib.request.urlretrieve(binary_url, binary_name)
    except Exception as e:
        print(f"Error downloading binary: {e}")
        return 1

    # Make binary executable and run it
    try:
        os.chmod(binary_name, 0o755)
        result = _run(f"./{binary_name} --lang python")
        return result
    except Exception as e:
        print(f"Error executing binary: {e}")
        return 1

if __name__ == "__main__":
    restack_get_started()
