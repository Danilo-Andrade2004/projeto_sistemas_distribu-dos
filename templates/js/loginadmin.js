const LABELS = {
  sala_de_aula:          'Sala de Aula',
  conversar_com_colegas: 'Colegas',
  professores:           'Professores',
  campus:                'Campus',
  emocional_semana:      'Emocional',
  motivacao_estudos:     'Motivação',
  ansiedade_escolar:     'Ansiedade',
  voz_na_escola:         'Representatividade',
  qualidade_sono:        'Sono',
  bem_estar_geral:       'Bem-estar',
};

const CAMPOS = Object.keys(LABELS);

function pillClass(val) {
  if (!val) return '';
  const v = val.toUpperCase();
  if (v === 'BEM' || v.includes('MUITO BEM')) return 'pill pill-bem';
  if (v === 'MAIS OU MENOS' || v === 'REGULAR') return 'pill pill-meio';
  return 'pill pill-mal';
}

function renderTabela(dados) {
  if (!dados || dados.length === 0) {
    document.getElementById('tabela-conteudo').innerHTML = `
      <div class="empty-state">
        <div class="icon">🔎</div>
        <p>Nenhum registro encontrado.</p>
      </div>`;
    document.getElementById('badge-total').textContent = '0 registros';
    return;
  }

  const arr = Array.isArray(dados) ? dados : [dados];
  document.getElementById('badge-total').textContent = `${arr.length} registro(s)`;

  const th = ['ID', 'Usuário', ...Object.values(LABELS), 'Comentário']
    .map(h => `<th>${h}</th>`).join('');

  const rows = arr.map(r => {
    const tds = CAMPOS.map(c =>
      `<td><span class="${pillClass(r[c])}">${r[c] || '-'}</span></td>`
    ).join('');
    const comentario = r.comentario
      ? `<td title="${r.comentario}">💬 ${r.comentario.substring(0, 30)}${r.comentario.length > 30 ? '...' : ''}</td>`
      : '<td>—</td>';
    return `<tr><td><strong>#${r.id}</strong></td><td>${r.usuario_id}</td>${tds}${comentario}</tr>`;
  }).join('');

  document.getElementById('tabela-conteudo').innerHTML = `
    <table>
      <thead><tr>${th}</tr></thead>
      <tbody>${rows}</tbody>
    </table>`;
}

function mostrarToast(msg, erro = false) {
  const t = document.getElementById('toast');
  t.textContent = msg;
  t.className = 'toast' + (erro ? ' erro' : '');
  t.classList.add('show');
  setTimeout(() => t.classList.remove('show'), 3000);
}

async function fazerLogin() {
  const email = document.getElementById('login-email').value.trim();
  const senha = document.getElementById('login-senha').value;
  const msgErro = document.getElementById('msg-erro-login');

  if (!email || !senha) {
    msgErro.textContent = '⚠️ Preencha e-mail e senha.';
    msgErro.style.display = 'block';
    return;
  }

  try {
    const res = await fetch('http://localhost:8080/loginadmin', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, senha })
    });

    if (!res.ok) {
      msgErro.style.display = 'block';
      msgErro.textContent = '❌ E-mail ou senha inválidos.';
      return;
    }

    msgErro.style.display = 'none';
    document.getElementById('tela-login').style.display = 'none';
    document.getElementById('tela-dashboard').style.display = 'block';
    mostrarToast('✅ Login realizado com sucesso!');

  } catch (e) {
    msgErro.textContent = '❌ Erro de conexão com o servidor.';
    msgErro.style.display = 'block';
  }
}

async function listarTodos() {
  document.getElementById('tabela-conteudo').innerHTML = '<div class="loading">⏳ Carregando respostas...</div>';

  try {
    const res = await fetch('http://localhost:8080/listarquestionario');
    const data = await res.json();
    if (!res.ok) { mostrarToast('Erro ao buscar dados.', true); return; }
    renderTabela(data);
  } catch (e) {
    mostrarToast('Erro de conexão.', true);
  }
}

async function buscarPorId() {
  const id = document.getElementById('input-id').value.trim();
  if (!id) { mostrarToast('⚠️ Digite um ID.', true); return; }

  document.getElementById('tabela-conteudo').innerHTML = '<div class="loading">⏳ Buscando...</div>';

  try {
    const res = await fetch(`http://localhost:8080/buscarquestionario?id=${id}`);

    if (res.status === 404) {
      document.getElementById('tabela-conteudo').innerHTML = `
        <div class="empty-state">
          <div class="icon">🔎</div>
          <p>Questionário #${id} não encontrado.</p>
        </div>`;
      document.getElementById('badge-total').textContent = '0 registros';
      return;
    }

    const data = await res.json();
    renderTabela(data);
  } catch (e) {
    mostrarToast('Erro de conexão.', true);
  }
}

function sair() {
  document.getElementById('tela-dashboard').style.display = 'none';
  document.getElementById('tela-login').style.display = 'flex';
  document.getElementById('login-email').value = '';
  document.getElementById('login-senha').value = '';
  document.getElementById('input-id').value = '';
  document.getElementById('tabela-conteudo').innerHTML = `
    <div class="empty-state">
      <div class="icon">📊</div>
      <p>Clique em "Listar Todos" para ver as respostas</p>
    </div>`;
}

// Enter no campo de senha faz login
document.getElementById('login-senha').addEventListener('keydown', e => {
  if (e.key === 'Enter') fazerLogin();
});

// Enter no campo de busca faz a busca
document.getElementById('input-id').addEventListener('keydown', e => {
  if (e.key === 'Enter') buscarPorId();
});
