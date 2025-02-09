document.body.addEventListener("showToast", function (e) {
  const toastPayload = e?.detail;
  const toast = document.getElementById("toast");

  if (toastPayload?.level === "ERROR") {
    toast.classList.add("toast-error");
  } else {
    toast.classList.add("toast-success");
  }

  const toastContent = createToastContent(
    toastPayload.title,
    toastPayload.message
  );

  toast.appendChild(toastContent);
  toast.style.display = "block";
  toast.classList.add("show-toast");

  setTimeout(() => {
    toast.remove();
  }, 6000);

});

function createToastContent(title, message) {
  const toastContent = document.createElement("div");
  toastContent.classList.add("toast-content");

  const toastTitle = document.createElement("h3");
  toastTitle.classList.add("toast-title");
  toastTitle.textContent = title;

  const toastMessage = document.createElement("p");
  toastMessage.classList.add("toast-message");
  toastMessage.textContent = message;

  toastContent.appendChild(toastTitle);
  toastContent.appendChild(toastMessage);

  return toastContent;
}
