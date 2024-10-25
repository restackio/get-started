const { execSync } = require('child_process');
const os = require('os');

function runBinary() {
  const arch = os.arch();
  const platform = os.platform();
  
  let binaryPath;
  
  if (platform === 'win32') {
    binaryPath = arch === 'x64' ? './binary_win64.exe' : './binary_win32.exe';
  } else if (platform === 'darwin') {
    binaryPath = './binary_mac';
  } else if (platform === 'linux') {
    binaryPath = arch === 'x64' ? './binary_linux64' : './binary_linux32';
  } else {
    throw new Error('Unsupported platform');
  }

  try {
    const output = execSync(binaryPath, { encoding: 'utf-8' });
    console.log(output);
  } catch (error) {
    console.error('Error executing binary:', error.message);
  }
}

runBinary();