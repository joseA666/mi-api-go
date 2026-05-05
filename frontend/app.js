const API = '/api';

const $ = id => document.getElementById(id);

const tabla        = $('tabla');
const tablaBody    = $('tabla-body');
const stateLoading = $('state-loading');
const stateEmpty   = $('state-empty');
const stateError   = $('state-error');
const productCount = $('product-count');
const formProducto = $('form-producto');
const formError    = $('form-error');
const modal        = $('modal');
const formEditar   = $('form-editar');
const editError    = $('edit-error');
const toast        = $('toast');

// ── Toast ──────────────────────────────────────────────────────────────────
let toastTimer;
function showToast(msg, type = 'ok') {
  clearTimeout(toastTimer);
  toast.textContent = msg;
  toast.className = `toast ${type}`;
  toastTimer = setTimeout(() => { toast.className = 'toast hidden'; }, 3000);
}

// ── Helpers ────────────────────────────────────────────────────────────────
function esc(str) {
  return String(str)
    .replace(/&/g, '&amp;').replace(/</g, '&lt;')
    .replace(/>/g, '&gt;').replace(/"/g, '&quot;');
}

function formatPrice(n) {
  return new Intl.NumberFormat('es-AR', { style: 'currency', currency: 'USD' }).format(n);
}

// ── Load products ──────────────────────────────────────────────────────────
async function cargarProductos() {
  stateLoading.classList.remove('hidden');
  tabla.classList.add('hidden');
  stateEmpty.classList.add('hidden');
  stateError.classList.add('hidden');
  productCount.textContent = '';

  try {
    const res = await fetch(`${API}/productos`);
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    const productos = await res.json();

    stateLoading.classList.add('hidden');

    if (!productos || productos.length === 0) {
      stateEmpty.classList.remove('hidden');
      return;
    }

    productCount.textContent =
      `${productos.length} producto${productos.length !== 1 ? 's' : ''}`;

    tablaBody.innerHTML = productos.map(p => `
      <tr>
        <td class="cell-id" title="${esc(p.id)}">${esc(p.id).slice(0, 8)}…</td>
        <td>${esc(p.nombre)}</td>
        <td class="cell-price">${formatPrice(p.precio)}</td>
        <td class="cell-actions">
          <button class="btn btn-sm btn-edit"
            data-action="editar"
            data-id="${esc(p.id)}"
            data-nombre="${esc(p.nombre)}"
            data-precio="${p.precio}">Editar</button>
          <button class="btn btn-sm btn-danger"
            data-action="eliminar"
            data-id="${esc(p.id)}"
            data-nombre="${esc(p.nombre)}">Eliminar</button>
        </td>
      </tr>
    `).join('');

    tabla.classList.remove('hidden');
  } catch {
    stateLoading.classList.add('hidden');
    stateError.textContent = 'No se pudo conectar con el servidor.';
    stateError.classList.remove('hidden');
  }
}

// ── Table click delegation ─────────────────────────────────────────────────
tablaBody.addEventListener('click', e => {
  const btn = e.target.closest('button[data-action]');
  if (!btn) return;
  const { action, id, nombre, precio } = btn.dataset;
  if (action === 'editar')   abrirEditar(id, nombre, parseFloat(precio));
  if (action === 'eliminar') eliminar(id, nombre);
});

// ── Create ─────────────────────────────────────────────────────────────────
formProducto.addEventListener('submit', async e => {
  e.preventDefault();
  formError.classList.add('hidden');

  const nombre = $('nombre').value.trim();
  const precio = parseFloat($('precio').value);

  try {
    const res = await fetch(`${API}/productos`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ nombre, precio }),
    });
    if (!res.ok) {
      const d = await res.json().catch(() => ({}));
      throw new Error(d.error || `HTTP ${res.status}`);
    }
    formProducto.reset();
    showToast('Producto agregado correctamente');
    cargarProductos();
  } catch (err) {
    formError.textContent = err.message;
    formError.classList.remove('hidden');
  }
});

// ── Edit modal ─────────────────────────────────────────────────────────────
function abrirEditar(id, nombre, precio) {
  $('edit-id').value    = id;
  $('edit-nombre').value = nombre;
  $('edit-precio').value = precio;
  editError.classList.add('hidden');
  modal.classList.remove('hidden');
  $('edit-nombre').focus();
}

$('btn-cancelar').addEventListener('click', () => modal.classList.add('hidden'));

modal.addEventListener('click', e => {
  if (e.target === modal) modal.classList.add('hidden');
});

document.addEventListener('keydown', e => {
  if (e.key === 'Escape') modal.classList.add('hidden');
});

formEditar.addEventListener('submit', async e => {
  e.preventDefault();
  editError.classList.add('hidden');

  const id     = $('edit-id').value;
  const nombre = $('edit-nombre').value.trim();
  const precio = parseFloat($('edit-precio').value);

  try {
    const res = await fetch(`${API}/productos/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ nombre, precio }),
    });
    if (!res.ok) {
      const d = await res.json().catch(() => ({}));
      throw new Error(d.error || `HTTP ${res.status}`);
    }
    modal.classList.add('hidden');
    showToast('Producto actualizado');
    cargarProductos();
  } catch (err) {
    editError.textContent = err.message;
    editError.classList.remove('hidden');
  }
});

// ── Delete ─────────────────────────────────────────────────────────────────
async function eliminar(id, nombre) {
  if (!confirm(`¿Eliminar "${nombre}"? Esta acción no se puede deshacer.`)) return;
  try {
    const res = await fetch(`${API}/productos/${id}`, { method: 'DELETE' });
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    showToast('Producto eliminado', 'err');
    cargarProductos();
  } catch {
    showToast('Error al eliminar el producto', 'err');
  }
}

// ── Init ───────────────────────────────────────────────────────────────────
cargarProductos();
