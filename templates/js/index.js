const selecionados = { curso: null, turno: null, sexo: null };

document.querySelectorAll('.opt-btn').forEach(btn => {
  btn.addEventListener('click', () => {
    const group = btn.dataset.group;
    document.querySelectorAll(`[data-group="${group}"]`).forEach(b => b.classList.remove('selected'));
    btn.classList.add('selected');
    selecionados[group] = parseInt(btn.dataset.value);
  });
});

function mostrarToast(msg, erro = false) {
  const t = document.getElementById('toast');
  t.textContent = msg;
  t.className = 'toast' + (erro ? ' erro' : '');
  t.classList.add('show');
  setTimeout(() => t.classList.remove('show'), 3000);
}

async function enviar() {
  const { curso, turno, sexo } = selecionados;
  if (!curso || !turno || !sexo) {
    document.getElementById('msg-erro').style.display = 'block';
    return;
  }
  document.getElementById('msg-erro').style.display = 'none';

  const btn = document.getElementById('btn-enviar');
  btn.disabled = true;
  btn.textContent = 'Enviando...';

  try {
    const res = await fetch('http://localhost:8080/cadastro', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ curso, turno, sexo })
    });

    const data = await res.json();

    if (!res.ok) {
      mostrarToast('Erro ao cadastrar. Tente novamente.', true);
      btn.disabled = false;
      btn.textContent = 'ENVIAR';
      return;
    }

    sessionStorage.setItem('usuario_id', data.usuario_id);
    window.location.href = 'questionario.html';

  } catch (e) {
    mostrarToast('Erro de conexão com o servidor.', true);
    btn.disabled = false;
    btn.textContent = 'ENVIAR';
  }
}
