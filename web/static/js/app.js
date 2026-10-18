/**
 * VOLTRA Electric Landing Page Application Logic
 * PT Mega Karya Perkasa
 */

document.addEventListener('DOMContentLoaded', () => {
  initProductFilter();
  initMobileMenu();
  initInquiryForm();
});

// Category Filter Handling
function initProductFilter() {
  const filterBtns = document.querySelectorAll('.filter-btn');
  const productCards = document.querySelectorAll('.product-card');

  filterBtns.forEach(btn => {
    btn.addEventListener('click', () => {
      // Toggle button active classes
      filterBtns.forEach(b => {
        b.classList.remove('bg-[#a8f020]', 'text-slate-950', 'shadow-sm');
        b.classList.add('text-slate-600');
      });
      btn.classList.add('bg-[#a8f020]', 'text-slate-950', 'shadow-sm');
      btn.classList.remove('text-slate-600');

      const filter = btn.getAttribute('data-filter');
      productCards.forEach(card => {
        const category = card.getAttribute('data-category');
        if (filter === 'all' || category === filter) {
          card.style.display = 'flex';
          card.classList.remove('opacity-0');
        } else {
          card.style.display = 'none';
        }
      });
    });
  });
}

// Modal open with API fetch
async function openProductModal(productId) {
  const modal = document.getElementById('product-modal');
  const content = document.getElementById('modal-content');
  
  if (!modal || !content) return;

  // Show loading skeleton first
  content.innerHTML = `
    <div class="p-8 text-center">
      <div class="inline-block w-8 h-8 border-4 border-lime-400 border-t-transparent rounded-full animate-spin mb-3"></div>
      <p class="text-xs text-slate-500 font-bold uppercase tracking-wider">Memuat Spesifikasi Trafo VOLTRA...</p>
    </div>
  `;
  modal.classList.remove('hidden');
  modal.classList.add('flex');

  try {
    const res = await fetch(`/api/products/${productId}`);
    const data = await res.json();

    if (data.status !== 'success' || !data.product) {
      throw new Error('Data tidak ditemukan');
    }

    const p = data.product;
    
    // Render detailed modal content
    content.innerHTML = `
      <div class="space-y-5">
        <div class="flex items-start justify-between">
          <div>
            <span class="inline-block bg-[#a8f020] text-slate-950 text-[10px] font-black uppercase px-2.5 py-1 rounded tracking-wider">
              ${p.series} • ${p.badge}
            </span>
            <h3 class="text-2xl font-black text-slate-900 mt-1.5">${p.name}</h3>
            <p class="text-xs font-bold text-sky-600">${p.subtitle}</p>
          </div>
        </div>

        <div class="relative h-44 rounded-2xl overflow-hidden bg-slate-900 flex items-center justify-center border border-slate-200">
          <img src="${p.image_url}" alt="${p.name}" class="w-full h-full object-cover">
          <div class="absolute inset-0 bg-gradient-to-t from-slate-950/80 via-transparent to-transparent"></div>
          <span class="absolute bottom-3 left-4 text-xs font-bold text-white bg-slate-900/80 px-3 py-1 rounded-md border border-slate-700">
            Kapasitas: ${p.capacity}
          </span>
        </div>

        <p class="text-sm text-slate-600 leading-relaxed">${p.description}</p>
        
        <div>
          <h4 class="font-bold text-xs uppercase tracking-wider text-slate-800 mb-2">Fitur Utama & Keunggulan:</h4>
          <ul class="text-xs text-slate-600 space-y-1.5 list-disc pl-4">
            ${p.features.map(f => `<li>${f}</li>`).join('')}
          </ul>
        </div>

        <div>
          <h4 class="font-bold text-xs uppercase tracking-wider text-slate-800 mb-2">Ringkasan Spesifikasi Teknis:</h4>
          <div class="grid grid-cols-2 gap-2 text-xs">
            ${Object.entries(p.specs || {}).map(([k, v]) => `
              <div class="bg-slate-50 p-2.5 rounded-xl border border-slate-200">
                <span class="text-slate-400 block text-[10px] uppercase font-bold">${k}</span>
                <span class="text-slate-800 font-bold">${v}</span>
              </div>
            `).join('')}
          </div>
        </div>

        <div class="pt-2">
          <a href="https://wa.me/18732768961?text=Halo%20PT%20Mega%20Karya%20Perkasa,%20saya%20tertarik%20dengan%20${encodeURIComponent(p.name)}%20(voltra-electric.id)" target="_blank" class="btn-lime w-full py-3.5 rounded-xl text-center text-xs font-black uppercase tracking-wider flex items-center justify-center gap-2">
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24"><path d="M.057 24l1.687-6.163c-1.041-1.804-1.588-3.849-1.587-5.946.003-6.556 5.338-11.891 11.893-11.891 3.181.001 6.167 1.24 8.413 3.488 2.245 2.248 3.481 5.236 3.48 8.414-.003 6.557-5.338 11.892-11.893 11.892-1.99-.001-3.951-.5-5.688-1.448l-6.305 1.654zm6.597-3.807c1.676.995 3.276 1.591 5.392 1.592 5.448 0 9.886-4.434 9.889-9.885.002-5.462-4.415-9.89-9.881-9.892-5.452 0-9.887 4.434-9.889 9.884-.001 2.225.651 3.891 1.746 5.634l-.999 3.648 3.742-.981z"/></svg>
            Minta Penawaran Resmi via WhatsApp
          </a>
        </div>
      </div>
    `;
  } catch (err) {
    content.innerHTML = `
      <div class="p-6 text-center text-rose-500">
        <p class="font-bold">Gagal memuat detail produk.</p>
        <p class="text-xs text-slate-500 mt-1">Silakan coba beberapa saat lagi atau hubungi sales langsung.</p>
      </div>
    `;
  }
}

function closeProductModal() {
  const modal = document.getElementById('product-modal');
  if (modal) {
    modal.classList.add('hidden');
    modal.classList.remove('flex');
  }
}

// Inquiry Form Handling
function initInquiryForm() {
  const form = document.getElementById('trafo-form');
  if (!form) return;

  const btn = document.getElementById('submit-inquiry-btn');
  if (!btn) return;

  btn.addEventListener('click', async (e) => {
    e.preventDefault();

    const clientName = document.getElementById('spec-client')?.value || '';
    const tankType = document.getElementById('spec-tank')?.value || '';
    const capacityKVA = document.getElementById('spec-kva')?.value || '';
    const winding = document.getElementById('spec-winding')?.value || '';
    const voltage = document.getElementById('spec-voltage')?.value || '';
    const notes = document.getElementById('spec-notes')?.value || '';

    try {
      const res = await fetch('/api/inquiry', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json'
        },
        body: JSON.stringify({
          client_name: clientName,
          tank_type: tankType,
          capacity_kva: capacityKVA,
          winding: winding,
          voltage: voltage,
          notes: notes
        })
      });

      const data = await res.json();
      if (data.status === 'success' && data.whatsapp_url) {
        window.open(data.whatsapp_url, '_blank');
      } else {
        alert('Gagal membuat tautan WhatsApp. Silakan gunakan tombol kontak langsung.');
      }
    } catch (err) {
      console.error(err);
      // Fallback direct URL creation
      const msg = `Halo PT Mega Karya Perkasa,\nSaya (${clientName || 'Pelanggan'}) ingin menanyakan Trafo VOLTRA:\n- Tangki: ${tankType}\n- Kapasitas: ${capacityKVA}\n- Bahan Belitan: ${winding}\n- Tegangan: ${voltage}`;
      window.open(`https://wa.me/18732768961?text=${encodeURIComponent(msg)}`, '_blank');
    }
  });
}

// Mobile Menu
function initMobileMenu() {
  const mobileBtn = document.getElementById('mobile-menu-btn');
  const mobileMenu = document.getElementById('mobile-menu');

  if (mobileBtn && mobileMenu) {
    mobileBtn.addEventListener('click', () => {
      mobileMenu.classList.toggle('hidden');
    });

    mobileMenu.querySelectorAll('a').forEach(link => {
      link.addEventListener('click', () => {
        mobileMenu.classList.add('hidden');
      });
    });
  }

  // Close modal when clicking outside
  window.addEventListener('click', (e) => {
    const modal = document.getElementById('product-modal');
    if (e.target === modal) {
      closeProductModal();
    }
  });
}

