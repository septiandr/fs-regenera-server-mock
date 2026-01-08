const { execSync } = require("child_process");

const machineName = "podman-machine-default";
const imageName = "fs-regenera-app";
const containerName = "fs-regenera-service";
const portMapping = "9070:9070";

try {
  console.log("Mengecek status Podman machine...");

  let status = "";
  try {
    // inspect machine, jika field Status ada ambil
    const out = execSync(
      `podman machine inspect ${machineName} --format '{{.Status}}'`
    )
      .toString()
      .trim();
    status = out || "Stopped"; // jika kosong, anggap stopped
  } catch {
    console.log(
      `Podman machine "${machineName}" sudah ada tapi gagal inspect. Lanjut tanpa init...`
    );
    status = "Stopped"; // jangan init lagi
  }

  if (status !== "Running") {
    console.log("Menjalankan Podman machine...");
    execSync(`podman machine start ${machineName}`, { stdio: "inherit" });
  } else {
    console.log("Podman machine sudah berjalan.");
  }

  // Build image
  console.log("Membangun image Podman...");
  execSync(`podman build -t ${imageName} .`, { stdio: "inherit" });

  // Jalankan container
  console.log("Menjalankan container Podman...");
  execSync(
    `podman run -d -p ${portMapping} --name ${containerName} ${imageName}`,
    { stdio: "inherit" }
  );

  console.log("Container berhasil dijalankan.");
} catch (err) {
  console.error("Terjadi error:", err.message);
}
