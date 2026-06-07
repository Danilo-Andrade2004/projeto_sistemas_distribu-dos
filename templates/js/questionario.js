const usuarioId = parseInt(sessionStorage.getItem('usuario_id'));
if (!usuarioId) window.location.href = 'index.html';

const PERGUNTAS = [
  { campo: 'sala_de_aula',          label: 'Como você se sente em sala de aula?' },
  { campo: 'conversar_com_colegas', label: 'Como você se sente ao conversar com seus colegas?' },
  { campo: 'professores',           label: 'Como você se sente em relação aos professores?' },
  { campo: 'campus',                label: 'Como você se sente no ambiente do campus?' },
  { campo: 'emocional_semana',      label: 'Como você avalia o seu emocional nesta semana?' },
  { campo: 'motivacao_estudos',     label: 'Como você avalia a sua motivação para estudar?' },
  { campo: 'ansiedade_escolar',     label: 'Como você se sente em relação à ansiedade escolar?' },
  { campo: 'voz_na_escola',         label: 'Qual é a sua avaliação em relação à representatividade dentro da escola?' },
  { campo: 'qualidade_sono',        label: 'Como você avalia a qualidade do seu sono?' },
  { campo: 'bem_estar_geral',       label: 'Como você avalia o seu bem-estar geral?' },
];

const EMOJIS = [
  { val: 1, face: '😄', label: 'Muito Bem' },
  { val: 2, face: '🙂', label: 'Bem' },
  { val: 3, face: '😐', label: 'Regular' },
  { val: 4, face: '😕', label: 'Mal' },
  { val: 5, face: '😢', label: 'Muito Mal' },
];

const respostas = {};

function renderPerguntas() {
  const container = document.getElementById('perguntas-container');
  PERGUNTAS.forEach((p, i) => {
    const card = document.createElement('div');
    card.className = 'pergunta-card';
    card.id = `card-${i}`;

    const emojisHTML = EMOJIS.map(e => `
      <button class="emoji-btn" data-campo="${p.campo}" data-val="${e.val}" onclick="selecionar('${p.campo}', ${e.val}, ${i})">
        <span class="emoji-face">${e.face}</span>
        <span class="emoji-label">${e.label}</span>
      </button>
    `).join('');

    card.innerHTML = `
      <span class="pergunta-titulo">${i + 1}. ${p.label}</span>
      <div class="emojis-row">${emojisHTML}</div>
    `;
    container.appendChild(card);
  });
}

function selecionar(campo, val, idx) {
  respostas[campo] = val;
  document.querySelectorAll(`[data-campo="${campo}"]`).forEach(b => b.classList.remove('selected'));
  document.querySelector(`[data-campo="${campo}"][data-val="${val}"]`).classList.add('selected');
  document.getElementById(`card-${idx}`).classList.add('respondida');
  atualizarProgresso();
}

function atualizarProgresso() {
  const total = PERGUNTAS.length;
  const respondidas = Object.keys(respostas).length;
  const pct = (respondidas / total) * 100;
  document.getElementById('progress-bar').style.width = pct + '%';
  document.getElementById('progress-label').textContent = `${respondidas} de ${total} respondidas`;
}

function mostrarToast(msg, tipo = '') {
  const t = document.getElementById('toast');
  t.textContent = msg;
  t.className = 'toast ' + tipo;
  t.classList.add('show');
  setTimeout(() => t.classList.remove('show'), 3500);
}

async function enviar() {
  const naoRespondidas = PERGUNTAS.filter(p => !respostas[p.campo]);
  if (naoRespondidas.length > 0) {
    mostrarToast(`⚠️ Responda todas as ${PERGUNTAS.length} perguntas antes de enviar.`, 'erro');
    return;
  }

  const btn = document.getElementById('btn-enviar');
  btn.disabled = true;
  btn.textContent = 'Enviando...';

  const payload = {
    usuario_id: usuarioId,
    quer_deixar_um_comentario: document.getElementById('comentario').value.trim(),
    ...respostas
  };

  try {
    const res = await fetch('http://localhost:8080/questionario', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    });

    const data = await res.json();

    if (!res.ok) {
      mostrarToast('Erro ao enviar. Tente novamente.', 'erro');
      btn.disabled = false;
      btn.textContent = 'ENVIAR RESPOSTAS';
      return;
    }

    sessionStorage.removeItem('usuario_id');
    document.getElementById('modal').classList.add('show');

  } catch (e) {
    mostrarToast('Erro de conexão com o servidor.', 'erro');
    btn.disabled = false;
    btn.textContent = 'ENVIAR RESPOSTAS';
  }
}

renderPerguntas();
