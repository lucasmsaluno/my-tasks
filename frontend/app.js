const modalBackdrop = document.querySelector(".modal-backdrop");
const addTaskButton = document.querySelector(".add-task");
const tasksArea = document.querySelector(".tasks");
const form = document.querySelector("form");
let editingID = null;

const formatDate = iso => { 
  const date = new Date(iso); 
  return date.toLocaleDateString("en-GB") + " at " + date.toLocaleTimeString("en-GB", { hour: "2-digit", minute: "2-digit" }); 
};

const openModal = () => modalBackdrop.classList.add("show");
const hideModal = () => modalBackdrop.classList.remove("show");
const closeModal = e => {
  if (e.target.classList.contains("modal-backdrop")) hideModal();
};

const loadTasks = async () => {
  try {
    const response = await fetch("http://localhost:8080/tasks");
    if (!response.ok) throw new Error("Failed to load tasks");
    
    const tasks = await response.json() || [];
    tasksArea.innerHTML = "";

    tasks.forEach(task => {
      const taskElement = document.createElement("div");
      taskElement.classList.add("task");
      taskElement.innerHTML = ` 
        <div class="task-card" data-card="${task.id}"> 
          <p>${task.title}</p> 
          <div style="display: flex; gap: 10px;">
            <button data-button="${task.id}">+</button> 
            <button class="delete-btn" style="background: transparent; font-size: 12px;" data-delete="${task.id}">❌</button> 
          </div>
        </div> 
        <div class="task-content"> 
          <div class="task-body">${task.description || ""}</div>
          <div class="content-bottom"> 
            <p>Last Update:</p> 
            <p>${formatDate(task.updated_at)}</p> 
          </div> 
        </div>`;
      tasksArea.appendChild(taskElement);
    });

  } catch (err) {
    console.error(err);
    alert("Error loading tasks");
  }
};

const createTask = async () => {
  const formTitle = document.querySelector(".title").value.trim();
  const formContent = document.querySelector(".content").value.trim();
  if (!formTitle) return alert("Task title is required.");

  try {
    const response = await fetch("http://localhost:8080/tasks", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: formTitle, description: formContent })
    });
    if (!response.ok) throw new Error("Error creating task");

    await loadTasks();
    hideModal();
    form.reset();
  } catch (err) {
    console.error(err);
    alert("Error creating task");
  }
};

const updateTask = async () => {
  const formTitle = document.querySelector(".title").value.trim();
  const formContent = document.querySelector(".content").value.trim();
  if (!formTitle) return alert("Task title is required.");

  try {
    const response = await fetch(`http://localhost:8080/tasks/${editingID}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: formTitle, description: formContent })
    });
    if (!response.ok) throw new Error("Error updating task");

    await loadTasks();
    hideModal();
    form.reset();
    editingID = null;
  } catch (err) {
    console.error(err);
    alert("Error updating task");
  }
};

const deleteTask = async (taskId) => {
  if (!confirm("Are you sure you want to delete this task?")) return;

  try {
    const response = await fetch(`http://localhost:8080/tasks/${taskId}`, {
      method: "DELETE"
    });
    if (!response.ok) throw new Error("Error deleting task");

    await loadTasks();
  } catch (err) {
    console.error(err);
    alert("Error deleting task");
  }
};

const tasksAction = e => {
  const clickedElement = e.target;

  // Delete button
  if (clickedElement.classList.contains("delete-btn")) {
    const taskId = clickedElement.dataset.delete;
    deleteTask(taskId);
    return;
  }

  if (clickedElement.tagName === "BUTTON" && clickedElement.dataset.button) {
    const taskCard = clickedElement.closest(".task");
    const taskContent = taskCard.querySelector(".task-content");
    taskContent.classList.toggle("show");
    return;
  }

  const taskCard = clickedElement.closest(".task-card");
  if (!taskCard) return;

  editingID = taskCard.dataset.card;
  openModal();

  const title = taskCard.querySelector("p").textContent;
  const content = taskCard.nextElementSibling.querySelector(".task-body").textContent.trim();

  document.querySelector(".title").value = title;
  document.querySelector(".content").value = content;
};

const submitForm = e => {
  e.preventDefault();
  editingID ? updateTask() : createTask();
};

tasksArea.addEventListener("click", tasksAction);
addTaskButton.addEventListener("click", openModal);
modalBackdrop.addEventListener("click", closeModal);
form.addEventListener("submit", submitForm);
window.addEventListener("DOMContentLoaded", loadTasks);
