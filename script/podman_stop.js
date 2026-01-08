const { execSync } = require("child_process");

const containerName = "fs-regenera-service";

try {
  console.log("Menghentikan container...");

  try {
    execSync(`podman stop ${containerName}`, { stdio: "inherit" });
  } catch {
    console.log(
      `Container "${containerName}" tidak ditemukan atau sudah berhenti. Lanjut...`
    );
  }

  console.log("Menghapus container...");
  try {
    execSync(`podman rm ${containerName}`, { stdio: "inherit" });
  } catch {
    console.log(
      `Container "${containerName}" tidak ditemukan atau sudah dihapus. Lanjut...`
    );
  }

  console.log("Mematikan Podman machine...");
  try {
    execSync(`podman machine stop`, { stdio: "inherit" });
  } catch (err) {
    console.log("Gagal mematikan Podman machine, kemungkinan sudah mati.");
  }

  console.log("Selesai.");
} catch (err) {
  console.error("Terjadi error:", err.message);
}
