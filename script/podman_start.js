const { execSync } = require("child_process");

const machineName = "podman-machine-default";
const imageName = "fs-regenera-app";
const containerName = "fs-regenera-service";
const portMapping = "9070:9070";

function run(cmd) {
  console.log(`\n$ ${cmd}`);
  execSync(cmd, { stdio: "inherit" });
}

function exec(cmd) {
  return execSync(cmd, { stdio: "pipe" }).toString().trim();
}

function exists(cmd) {
  try {
    execSync(cmd, { stdio: "ignore" });
    return true;
  } catch {
    return false;
  }
}

try {
  console.log("🔍 Cek Podman machine...");

  if (!exists(`podman machine inspect ${machineName}`)) {
    throw new Error(`Podman machine "${machineName}" belum ada`);
  }

  const inspectRaw = exec(`podman machine inspect ${machineName}`);
  const inspect = JSON.parse(inspectRaw)[0];

  const isRunning = inspect.State === "running" || inspect.Running === true;

  if (!isRunning) {
    console.log("▶️ Menjalankan Podman machine...");
    run(`podman machine start ${machineName}`);
  } else {
    console.log("✅ Podman machine sudah running");
  }

  console.log("🔁 Build image...");
  run(`podman build -t ${imageName} .`);

  console.log("🧹 Stop & hapus container lama (jika ada)...");
  if (exists(`podman container exists ${containerName}`)) {
    run(`podman stop ${containerName}`);
    run(`podman rm ${containerName}`);
  } else {
    console.log("ℹ️ Container belum ada, skip");
  }

  console.log("🚀 Menjalankan container...");
  run(`podman run -d -p ${portMapping} --name ${containerName} ${imageName}`);

  console.log("\n✅ Build & run berhasil!");
} catch (err) {
  console.error("\n❌ Error:", err.message);
  process.exit(1);
}
