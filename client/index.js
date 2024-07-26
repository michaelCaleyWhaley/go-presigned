const formSubmit = document.getElementById("form-submit");
const formFile = document.getElementById("form-file");

formSubmit?.addEventListener("click", async (e) => {
  e.preventDefault();
  const fileName = formFile?.files?.[0]?.name?.toLowerCase?.();
  if (typeof fileName !== "string") return;

  const response = await fetch(`/presigned`, {
    method: "POST",
    body: JSON.stringify({ fileName }),
  }).catch(() => null);
  if (!response) return;

  const { presignedUrl } = await response.json().catch(() => ({}));

  const formData = new FormData(formSubmit?.parentElement, formSubmit);
  const fileUpload = formData.get("file-upload");

  await fetch(presignedUrl, {
    method: "PUT",
    body: fileUpload,
  }).catch((err) => {
    console.log("err: ", err);
  });
});
