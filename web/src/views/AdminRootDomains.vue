<template>
  <div class="container mx-auto px-4 sm:px-6 lg:px-8 py-8 max-w-7xl space-y-6">
    <!-- Header -->
    <div class="flex justify-between items-center">
      <div>
        <h1 class="text-3xl font-bold">{{ $t('admin.rootDomainManagement') }}</h1>
        <p class="text-lg opacity-70 mt-2">{{ $t('admin.rootDomainManagementDesc') }}</p>
      </div>
      <button @click="showCreateModal = true" class="btn btn-primary gap-2">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        {{ $t('admin.addRootDomain') }}
      </button>
    </div>

    <!-- Root Domains List -->
    <div v-if="loading" class="flex justify-center py-12">
      <span class="loading loading-spinner loading-lg"></span>
    </div>

    <div v-else-if="rootDomains.length === 0" class="card bg-base-200 shadow-xl">
      <div class="card-body items-center text-center py-12">
        <div class="mb-6">
          <div class="w-24 h-24 rounded-full bg-base-300 flex items-center justify-center mx-auto">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 text-base-content opacity-40" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9" />
            </svg>
          </div>
        </div>
        <h3 class="card-title text-2xl mb-2">{{ $t('admin.noRootDomains') }}</h3>
        <p class="mb-4 opacity-70">{{ $t('admin.noRootDomainsDesc') }}</p>
        <button @click="showCreateModal = true" class="btn btn-primary gap-2">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          {{ $t('admin.addRootDomain') }}
        </button>
      </div>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="domain in rootDomains"
        :key="domain.id"
        class="card bg-base-100 shadow-xl border border-base-300 hover:shadow-2xl transition-all duration-300"
        :class="{ 'opacity-60': !domain.is_active }"
      >
        <div class="card-body">
          <div class="flex justify-between items-start mb-2">
            <div class="flex items-center gap-2 flex-wrap">
              <h2 class="card-title text-2xl font-mono">.{{ domain.domain }}</h2>
              <div class="flex gap-1.5">
                <span v-if="domain.is_hot" class="px-2 py-0.5 rounded-full text-xs font-medium inline-flex items-center gap-1 bg-warning/10 text-warning">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 18.657A8 8 0 016.343 7.343S7 9 9 10c0-2 .5-5 2.986-7C14 5 16.09 5.777 17.656 7.343A7.975 7.975 0 0120 13a7.975 7.975 0 01-2.343 5.657z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.879 16.121A3 3 0 1012.015 11L11 14H9c0 .768.293 1.536.879 2.121z" />
                  </svg>
                  {{ $t('admin.hot') }}
                </span>
                <span v-if="domain.is_new" class="px-2 py-0.5 rounded-full text-xs font-medium inline-flex items-center gap-1 bg-info/10 text-info">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
                  </svg>
                  {{ $t('admin.new') }}
                </span>
              </div>
            </div>
            <div class="flex gap-2">
              <button @click="editDomain(domain)" class="btn btn-ghost btn-sm btn-square" title="Edit">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                </svg>
              </button>
              <button @click="deleteDomain(domain)" class="btn btn-ghost btn-sm btn-square text-error" title="Delete">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
              </button>
            </div>
          </div>

          <p v-if="domain.description" class="text-sm opacity-70 mb-4">{{ domain.description }}</p>

          <div class="divider my-2"></div>

          <div class="space-y-2">
            <div class="flex justify-between items-center">
              <span class="text-sm font-medium opacity-70">{{ $t('admin.status') }}:</span>
              <span class="px-3 py-1 rounded-full text-xs font-medium inline-flex items-center gap-1.5" :class="domain.is_active ? 'bg-success/10 text-success' : 'bg-error/10 text-error'">
                <svg v-if="domain.is_active" xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                {{ domain.is_active ? $t('admin.active') : $t('admin.inactive') }}
              </span>
            </div>

            <div class="flex justify-between items-center">
              <span class="text-sm font-medium opacity-70">Priority:</span>
              <span class="px-3 py-1 rounded-full text-xs font-medium inline-flex items-center gap-1.5 bg-primary/10 text-primary">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11l5-5m0 0l5 5m-5-5v12" />
                </svg>
                {{ domain.priority }}
              </span>
            </div>

            <div class="flex justify-between items-center">
              <span class="text-sm font-medium opacity-70">Registrations:</span>
              <router-link
                :to="`/admin/root-domains/${domain.id}/domains`"
                class="px-3 py-1 rounded-full text-xs font-medium inline-flex items-center gap-1.5 bg-base-300 text-base-content hover:bg-primary/10 hover:text-primary transition-colors cursor-pointer"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
                {{ domain.registration_count || 0 }}
              </router-link>
            </div>

            <div class="flex justify-between items-center">
              <span class="text-sm font-medium opacity-70">DNSSEC:</span>
              <button
                v-if="domain.use_default_nameservers !== false"
                @click="openDNSSEC(domain)"
                class="px-3 py-1 rounded-full text-xs font-medium inline-flex items-center gap-1.5 transition-colors cursor-pointer"
                :class="dnssecBadgeClass(domain.id)"
                title="Manage DNSSEC"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
                {{ dnssecBadgeText(domain.id) }}
              </button>
              <span v-else class="text-xs opacity-60">External NS</span>
            </div>

            <div class="flex justify-between items-center">
              <span class="text-sm font-medium opacity-70">Pricing:</span>
              <span v-if="domain.is_free" class="px-3 py-1 rounded-full text-xs font-medium inline-flex items-center gap-1.5 bg-success/10 text-success">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                Free
              </span>
              <div v-else class="text-right">
                <div v-if="domain.price_per_year" class="text-xs font-medium text-primary">
                  {{ formatPrice(domain.price_per_year) }}/year
                </div>
                <div v-if="domain.lifetime_price" class="text-xs font-medium text-secondary">
                  {{ formatPrice(domain.lifetime_price) }} lifetime
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Create/Edit Modal -->
    <dialog :class="{ 'modal': true, 'modal-open': showCreateModal || showEditModal }">
      <div class="modal-box w-11/12 max-w-2xl">
        <h3 class="font-bold text-2xl mb-4">
          {{ showEditModal ? 'Edit Root Domain' : 'Add New Root Domain' }}
        </h3>

        <form @submit.prevent="showEditModal ? updateDomain() : createDomain()" class="space-y-4">
          <div class="form-control">
            <label class="label">
              <span class="label-text font-semibold">Domain</span>
            </label>
            <div class="relative">
              <input
                v-model="formData.domain"
                type="text"
                placeholder="example.com"
                class="input input-bordered w-full pl-8 font-mono"
                :disabled="showEditModal"
                required
              />
              <div class="absolute left-3 top-1/2 -translate-y-1/2 text-base-content font-mono font-bold text-lg">
                .
              </div>
            </div>
          </div>

          <div class="form-control">
            <label class="label">
              <span class="label-text font-semibold">Description (Optional)</span>
            </label>
            <textarea
              v-model="formData.description"
              class="textarea textarea-bordered h-20"
              placeholder="Brief description of this domain..."
            ></textarea>
          </div>

          <div class="form-control">
            <label class="label">
              <span class="label-text font-semibold">Priority</span>
            </label>
            <input
              v-model.number="formData.priority"
              type="number"
              class="input input-bordered"
              min="0"
              max="100"
            />
            <label class="label">
              <span class="label-text-alt opacity-60">Higher priority shows first (0-100)</span>
            </label>
          </div>

          <div class="divider">Pricing Configuration</div>

          <div class="form-control">
            <label class="label cursor-pointer justify-start gap-3">
              <input
                v-model="formData.is_free"
                type="checkbox"
                class="toggle toggle-success"
              />
              <span class="label-text font-semibold">{{ formData.is_free ? 'Free Domain' : 'Paid Domain' }}</span>
            </label>
            <label class="label">
              <span class="label-text-alt opacity-60">Free domains use user quota, paid domains require payment</span>
            </label>
          </div>

          <div v-if="!formData.is_free" class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div class="form-control">
              <label class="label">
                <span class="label-text font-semibold">Price Per Year ({{ currencySymbol }})</span>
              </label>
              <input
                v-model.number="formData.price_per_year"
                type="number"
                step="0.01"
                min="0"
                class="input input-bordered"
                placeholder="0.00"
              />
              <label class="label">
                <span class="label-text-alt opacity-60">Annual registration price</span>
              </label>
            </div>

            <div class="form-control">
              <label class="label">
                <span class="label-text font-semibold">Lifetime Price ({{ currencySymbol }})</span>
              </label>
              <input
                v-model.number="formData.lifetime_price"
                type="number"
                step="0.01"
                min="0"
                class="input input-bordered"
                placeholder="0.00"
              />
              <label class="label">
                <span class="label-text-alt opacity-60">One-time permanent registration price</span>
              </label>
            </div>
          </div>

          <div class="divider">Nameservers Configuration</div>

          <div class="form-control">
            <label class="label cursor-pointer justify-start gap-3">
              <input
                v-model="formData.use_default_nameservers"
                type="checkbox"
                class="toggle toggle-primary"
              />
              <span class="label-text font-semibold">{{ formData.use_default_nameservers ? 'Using Default Nameservers' : 'Using Custom Nameservers' }}</span>
            </label>
            <label class="label">
              <span class="label-text-alt opacity-60">Default nameservers are managed by the system</span>
            </label>
          </div>

          <div v-if="formData.use_default_nameservers" class="alert alert-info">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <div>
              <div class="font-bold">Default Nameservers (Read-only)</div>
              <div class="text-sm">ns1.nodelook.com, ns2.nodelook.com</div>
            </div>
          </div>

          <div v-else class="space-y-2">
            <div v-for="(ns, index) in formData.nameservers" :key="index" class="flex gap-2">
              <input
                v-model="formData.nameservers[index]"
                type="text"
                class="input input-bordered flex-1"
                placeholder="ns1.example.com"
              />
              <button
                v-if="formData.nameservers.length > 1"
                type="button"
                class="btn btn-ghost btn-square"
                @click="formData.nameservers.splice(index, 1)"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
            <button
              type="button"
              class="btn btn-sm btn-ghost gap-2"
              @click="formData.nameservers.push('')"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
              Add Nameserver
            </button>
          </div>

          <div class="divider">Display Options</div>

          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div class="form-control">
              <label class="label">
                <span class="label-text font-semibold">Status</span>
              </label>
              <label class="label cursor-pointer justify-start gap-3">
                <input
                  v-model="formData.is_active"
                  type="checkbox"
                  class="toggle toggle-success"
                />
                <span class="label-text">{{ formData.is_active ? 'Active' : 'Inactive' }}</span>
              </label>
            </div>

            <div class="form-control">
              <label class="label">
                <span class="label-text font-semibold">Min User Level</span>
              </label>
              <select v-model="formData.min_user_level" class="select select-bordered w-full">
                <option value="normal">Normal</option>
                <option value="basic">Basic</option>
                <option value="member">Member</option>
                <option value="regular">Regular</option>
                <option value="leader">Leader</option>
              </select>
              <label class="label">
                <span class="label-text-alt opacity-60">Only users at this level or above can register</span>
              </label>
            </div>

            <div class="form-control">
              <label class="label">
                <span class="label-text font-semibold">Mark as Hot</span>
              </label>
              <label class="label cursor-pointer justify-start gap-3">
                <input
                  v-model="formData.is_hot"
                  type="checkbox"
                  class="toggle toggle-warning"
                />
                <span class="label-text">{{ formData.is_hot ? 'Hot' : 'Normal' }}</span>
              </label>
            </div>

            <div class="form-control">
              <label class="label">
                <span class="label-text font-semibold">Mark as New</span>
              </label>
              <label class="label cursor-pointer justify-start gap-3">
                <input
                  v-model="formData.is_new"
                  type="checkbox"
                  class="toggle toggle-info"
                />
                <span class="label-text">{{ formData.is_new ? 'New' : 'Regular' }}</span>
              </label>
            </div>
          </div>

          <div class="modal-action">
            <button type="button" class="btn btn-ghost" @click="closeModals">Cancel</button>
            <button type="submit" class="btn btn-primary" :disabled="submitting">
              <span v-if="submitting" class="loading loading-spinner loading-sm"></span>
              <span v-else>{{ showEditModal ? 'Update' : 'Create' }}</span>
            </button>
          </div>
        </form>
      </div>
      <form method="dialog" class="modal-backdrop">
        <button type="button" @click="closeModals">close</button>
      </form>
    </dialog>

    <!-- DNSSEC Modal -->
    <dialog :class="{ 'modal': true, 'modal-open': showDNSSECModal }">
      <div class="modal-box w-11/12 max-w-2xl">
        <h3 class="font-bold text-2xl mb-1 font-mono">DNSSEC · .{{ dnssecDomain?.domain }}</h3>
        <p class="text-sm opacity-70 mb-4">
          The root zone must be signed and its DS record published at the registrar before any subdomain DNSSEC can validate.
        </p>

        <div v-if="dnssecState.loading" class="flex justify-center py-8">
          <span class="loading loading-spinner loading-lg"></span>
        </div>

        <div v-else class="space-y-4">
          <div class="flex items-center justify-between">
            <span class="badge" :class="dnssecState.enabled ? 'badge-success' : 'badge-ghost'">
              {{ dnssecState.enabled ? 'Signed' : 'Not signed' }}
            </span>
            <button
              @click="toggleRootDNSSEC"
              class="btn btn-sm"
              :class="dnssecState.enabled ? 'btn-error' : 'btn-success'"
              :disabled="dnssecState.submitting"
            >
              <span v-if="dnssecState.submitting" class="loading loading-spinner loading-xs"></span>
              <span v-else>{{ dnssecState.enabled ? 'Disable DNSSEC' : 'Enable DNSSEC' }}</span>
            </button>
          </div>

          <div v-if="!dnssecState.zoneExists" class="alert alert-warning text-sm">
            <span>Zone does not exist in PowerDNS yet. Enabling DNSSEC will create it with the configured nameservers.</span>
          </div>

          <div v-if="dnssecState.enabled" class="space-y-3">
            <div class="alert alert-info text-sm">
              <span>Add the DS record(s) below at your registrar for <strong>.{{ dnssecDomain?.domain }}</strong>. Remove them there before disabling DNSSEC.</span>
            </div>
            <div v-if="dnssecState.dsRecords.length === 0" class="text-sm opacity-60">No active DS records found.</div>
            <div v-for="(ds, idx) in dnssecState.dsRecords" :key="idx" class="flex items-center gap-2">
              <code class="bg-base-300 px-3 py-2 rounded text-xs font-mono break-all flex-1">{{ ds }}</code>
              <button @click="copyText(ds)" class="btn btn-ghost btn-xs btn-square flex-shrink-0" title="Copy">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                </svg>
              </button>
            </div>
            <div v-for="key in dnssecState.keys" :key="key.id">
              <div class="text-xs font-semibold uppercase tracking-wide opacity-50 mb-1">DNSKEY ({{ key.keytype.toUpperCase() }}, {{ key.algorithm }})</div>
              <div class="flex items-center gap-2">
                <code class="bg-base-300 px-3 py-2 rounded text-xs font-mono break-all flex-1">{{ key.dnskey }}</code>
                <button @click="copyText(key.dnskey)" class="btn btn-ghost btn-xs btn-square flex-shrink-0" title="Copy">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>

        <div v-if="!dnssecState.loading" class="divider my-2">Subdomain delegations</div>
        <div v-if="!dnssecState.loading" class="space-y-2">
          <p class="text-sm opacity-70">
            Publishes NS (and DS for signed subdomains) records for every subdomain into this zone. Run once after signing the root zone, or whenever subdomains show "chain of trust incomplete".
          </p>
          <div class="flex items-center gap-3">
            <button @click="repairDelegations" class="btn btn-sm btn-outline" :disabled="dnssecState.repairing || dnssecState.repair?.running">
              <span v-if="dnssecState.repairing || dnssecState.repair?.running" class="loading loading-spinner loading-xs"></span>
              <span v-else>Repair subdomain delegations</span>
            </button>
            <span v-if="dnssecState.repair" class="text-xs opacity-70">
              <template v-if="dnssecState.repair.running">Running… {{ dnssecState.repair.total }} domains</template>
              <template v-else>
                Done: {{ dnssecState.repair.delegated }} delegated ({{ dnssecState.repair.with_ds }} with DS), {{ dnssecState.repair.no_zone }} without zone, {{ dnssecState.repair.errors?.length || 0 }} errors
              </template>
            </span>
          </div>
          <div v-if="dnssecState.repair && !dnssecState.repair.running && dnssecState.repair.errors?.length" class="bg-base-200 rounded p-2 max-h-40 overflow-auto">
            <div v-for="(e, i) in dnssecState.repair.errors" :key="i" class="text-xs font-mono text-error break-all">{{ e }}</div>
          </div>
        </div>

        <div v-if="!dnssecState.loading" class="divider my-2">Orphan zones</div>
        <div v-if="!dnssecState.loading" class="space-y-2">
          <p class="text-sm opacity-70">
            Zones still present in PowerDNS whose domain no longer exists in the database. They keep answering (signed ones even SERVFAIL). Scan first, then delete them together with their NS/DS delegation.
          </p>
          <div class="flex items-center gap-3 flex-wrap">
            <button @click="scanOrphans" class="btn btn-sm btn-outline" :disabled="orphans.scanning">
              <span v-if="orphans.scanning" class="loading loading-spinner loading-xs"></span>
              <span v-else>Scan orphan zones</span>
            </button>
            <button
              v-if="orphans.zones.length > 0"
              @click="cleanupOrphans"
              class="btn btn-sm btn-error btn-outline"
              :disabled="orphans.cleaning || dnssecState.orphanCleanup?.running"
            >
              <span v-if="orphans.cleaning || dnssecState.orphanCleanup?.running" class="loading loading-spinner loading-xs"></span>
              <span v-else>Delete {{ orphans.zones.length }} orphan zones</span>
            </button>
            <span v-if="orphans.scanned" class="text-xs opacity-70">{{ orphans.zones.length }} orphan zone(s) found</span>
            <span v-if="dnssecState.orphanCleanup" class="text-xs opacity-70">
              <template v-if="dnssecState.orphanCleanup.running">Cleaning… {{ dnssecState.orphanCleanup.deleted }}/{{ dnssecState.orphanCleanup.total }}</template>
              <template v-else>Cleanup done: {{ dnssecState.orphanCleanup.deleted }}/{{ dnssecState.orphanCleanup.total }} deleted, {{ dnssecState.orphanCleanup.errors?.length || 0 }} errors</template>
            </span>
          </div>
          <div v-if="orphans.zones.length > 0" class="bg-base-200 rounded p-2 max-h-32 overflow-auto text-xs font-mono">
            {{ orphans.zones.slice(0, 200).join(', ') }}<span v-if="orphans.zones.length > 200"> … (+{{ orphans.zones.length - 200 }} more)</span>
          </div>
          <div v-if="dnssecState.orphanCleanup && !dnssecState.orphanCleanup.running && dnssecState.orphanCleanup.errors?.length" class="bg-base-200 rounded p-2 max-h-32 overflow-auto">
            <div v-for="(e, i) in dnssecState.orphanCleanup.errors" :key="i" class="text-xs font-mono text-error break-all">{{ e }}</div>
          </div>
        </div>

        <div class="modal-action">
          <button type="button" @click="closeDNSSEC" class="btn">Close</button>
        </div>
      </div>
      <form method="dialog" class="modal-backdrop">
        <button type="button" @click="closeDNSSEC">close</button>
      </form>
    </dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from '../utils/axios'
import { useToast } from '../composables/useToast'
import { useCurrency } from '../composables/useCurrency'

const toast = useToast()
const { formatPrice, currencySymbol } = useCurrency()
const loading = ref(true)
const submitting = ref(false)
const rootDomains = ref([])
const showCreateModal = ref(false)
const showEditModal = ref(false)
const editingDomain = ref(null)

const formData = ref({
  domain: '',
  description: '',
  priority: 0,
  is_active: true,
  is_hot: false,
  is_new: false,
  is_free: true,
  min_user_level: 'normal',
  price_per_year: null,
  lifetime_price: null,
  use_default_nameservers: true,
  nameservers: [''],
})

// DNSSEC state per root domain (id -> status) plus the modal state
const dnssecStatuses = ref({})
const showDNSSECModal = ref(false)
const dnssecDomain = ref(null)
const dnssecState = ref({
  loading: false,
  submitting: false,
  repairing: false,
  repair: null,
  orphanCleanup: null,
  zoneExists: false,
  enabled: false,
  keys: [],
  dsRecords: [],
})
let repairPollTimer = null

const pollRepair = (domainId) => {
  clearTimeout(repairPollTimer)
  repairPollTimer = setTimeout(async () => {
    if (!showDNSSECModal.value || dnssecDomain.value?.id !== domainId) return
    try {
      const response = await axios.get(`/api/admin/root-domains/${domainId}/dnssec`)
      applyRootDNSSEC(response.data)
      if (response.data.repair?.running || response.data.orphan_cleanup?.running) pollRepair(domainId)
    } catch (error) {
      console.error('Failed to poll repair status:', error)
    }
  }, 3000)
}

const orphans = ref({ scanning: false, cleaning: false, scanned: false, zones: [] })

const scanOrphans = async () => {
  const domain = dnssecDomain.value
  if (!domain) return
  orphans.value.scanning = true
  try {
    const response = await axios.get(`/api/admin/root-domains/${domain.id}/orphan-zones`)
    orphans.value.zones = response.data.zones || []
    orphans.value.scanned = true
  } catch (error) {
    toast.error(error.response?.data?.error || 'Failed to scan orphan zones')
  } finally {
    orphans.value.scanning = false
  }
}

const cleanupOrphans = async () => {
  const domain = dnssecDomain.value
  if (!domain) return
  if (!confirm(`Delete ${orphans.value.zones.length} orphan zone(s) under .${domain.domain} from PowerDNS?\n\nThese names have no domain in the database. This cannot be undone.`)) return
  orphans.value.cleaning = true
  try {
    const response = await axios.post(`/api/admin/root-domains/${domain.id}/orphan-zones/cleanup`)
    dnssecState.value.orphanCleanup = response.data.orphan_cleanup
    toast.success(response.data.message || 'Cleanup started')
    orphans.value.zones = []
    orphans.value.scanned = false
    pollRepair(domain.id)
  } catch (error) {
    if (error.response?.data?.orphan_cleanup) dnssecState.value.orphanCleanup = error.response.data.orphan_cleanup
    toast.error(error.response?.data?.error || 'Failed to start cleanup')
  } finally {
    orphans.value.cleaning = false
  }
}

const repairDelegations = async () => {
  const domain = dnssecDomain.value
  if (!domain) return
  if (!confirm(`Publish NS/DS delegations for all subdomains of .${domain.domain} into its zone?`)) return
  dnssecState.value.repairing = true
  try {
    const response = await axios.post(`/api/admin/root-domains/${domain.id}/dnssec/repair`)
    dnssecState.value.repair = response.data.repair
    toast.success(response.data.message || 'Repair started')
    pollRepair(domain.id)
  } catch (error) {
    if (error.response?.data?.repair) dnssecState.value.repair = error.response.data.repair
    toast.error(error.response?.data?.error || 'Failed to start repair')
  } finally {
    dnssecState.value.repairing = false
  }
}

onMounted(async () => {
  await fetchRootDomains()
})

const fetchRootDomains = async () => {
  loading.value = true
  try {
    const response = await axios.get('/api/admin/root-domains')
    rootDomains.value = response.data.root_domains || []
    fetchDNSSECStatuses()
  } catch (error) {
    console.error('Failed to fetch root domains:', error)
    toast.error(error.response?.data?.error || 'Failed to fetch root domains')
  } finally {
    loading.value = false
  }
}

const createDomain = async () => {
  submitting.value = true
  try {
    // Filter out empty nameservers
    const nameservers = formData.value.nameservers.filter(ns => ns.trim() !== '')

    const payload = {
      ...formData.value,
      nameservers: formData.value.use_default_nameservers ? [] : nameservers,
    }

    await axios.post('/api/admin/root-domains', payload)
    toast.success('Root domain created successfully!')
    closeModals()
    await fetchRootDomains()
  } catch (error) {
    toast.error(error.response?.data?.error || 'Failed to create root domain')
  } finally {
    submitting.value = false
  }
}

const editDomain = (domain) => {
  editingDomain.value = domain

  // Parse nameservers from JSON string
  let nameservers = ['']
  try {
    if (domain.nameservers) {
      nameservers = JSON.parse(domain.nameservers)
      if (!Array.isArray(nameservers) || nameservers.length === 0) {
        nameservers = ['']
      }
    }
  } catch (e) {
    console.error('Failed to parse nameservers:', e)
    nameservers = ['']
  }

  formData.value = {
    domain: domain.domain,
    description: domain.description || '',
    priority: domain.priority,
    is_active: domain.is_active,
    is_hot: domain.is_hot,
    is_new: domain.is_new,
    is_free: domain.is_free ?? true,
    min_user_level: domain.min_user_level || 'normal',
    price_per_year: domain.price_per_year || null,
    lifetime_price: domain.lifetime_price || null,
    use_default_nameservers: domain.use_default_nameservers ?? true,
    nameservers: nameservers,
  }
  showEditModal.value = true
}

const updateDomain = async () => {
  submitting.value = true
  try {
    // Filter out empty nameservers
    const nameservers = formData.value.nameservers.filter(ns => ns.trim() !== '')

    await axios.put(`/api/admin/root-domains/${editingDomain.value.id}`, {
      description: formData.value.description || null,
      priority: formData.value.priority,
      is_active: formData.value.is_active,
      is_hot: formData.value.is_hot,
      is_new: formData.value.is_new,
      is_free: formData.value.is_free,
      min_user_level: formData.value.min_user_level || 'normal',
      price_per_year: formData.value.is_free ? null : (formData.value.price_per_year || null),
      lifetime_price: formData.value.is_free ? null : (formData.value.lifetime_price || null),
      use_default_nameservers: formData.value.use_default_nameservers,
      nameservers: formData.value.use_default_nameservers ? [] : nameservers,
    })
    toast.success('Root domain updated successfully!')
    closeModals()
    await fetchRootDomains()
  } catch (error) {
    toast.error(error.response?.data?.error || 'Failed to update root domain')
  } finally {
    submitting.value = false
  }
}

const deleteDomain = async (domain) => {
  if (!confirm(`Are you sure you want to delete .${domain.domain}?\n\nThis cannot be undone.`)) {
    return
  }

  try {
    await axios.delete(`/api/admin/root-domains/${domain.id}`)
    toast.success('Root domain deleted successfully!')
    await fetchRootDomains()
  } catch (error) {
    toast.error(error.response?.data?.error || 'Failed to delete root domain')
  }
}

const fetchDNSSECStatuses = async () => {
  await Promise.all(
    rootDomains.value
      .filter(d => d.use_default_nameservers !== false)
      .map(async (d) => {
        try {
          const response = await axios.get(`/api/admin/root-domains/${d.id}/dnssec`)
          dnssecStatuses.value[d.id] = response.data
        } catch (error) {
          console.error(`Failed to fetch DNSSEC status for ${d.domain}:`, error)
        }
      })
  )
}

const dnssecBadgeClass = (id) => {
  const status = dnssecStatuses.value[id]
  if (!status) return 'bg-base-300 text-base-content'
  return status.enabled ? 'bg-success/10 text-success hover:bg-success/20' : 'bg-warning/10 text-warning hover:bg-warning/20'
}

const dnssecBadgeText = (id) => {
  const status = dnssecStatuses.value[id]
  if (!status) return 'Unknown'
  return status.enabled ? 'Signed' : 'Not signed'
}

const applyRootDNSSEC = (data) => {
  dnssecState.value.zoneExists = !!data.zone_exists
  dnssecState.value.enabled = !!data.enabled
  dnssecState.value.keys = data.keys || []
  dnssecState.value.dsRecords = data.ds_records || []
  if (data.repair !== undefined) dnssecState.value.repair = data.repair
  if (data.orphan_cleanup !== undefined) dnssecState.value.orphanCleanup = data.orphan_cleanup
  if (dnssecDomain.value) {
    dnssecStatuses.value[dnssecDomain.value.id] = data
  }
}

const openDNSSEC = async (domain) => {
  dnssecDomain.value = domain
  showDNSSECModal.value = true
  dnssecState.value.loading = true
  dnssecState.value.repair = null
  dnssecState.value.orphanCleanup = null
  orphans.value = { scanning: false, cleaning: false, scanned: false, zones: [] }
  try {
    const response = await axios.get(`/api/admin/root-domains/${domain.id}/dnssec`)
    applyRootDNSSEC(response.data)
    if (response.data.repair?.running || response.data.orphan_cleanup?.running) pollRepair(domain.id)
  } catch (error) {
    toast.error(error.response?.data?.error || 'Failed to fetch DNSSEC status')
  } finally {
    dnssecState.value.loading = false
  }
}

const closeDNSSEC = () => {
  clearTimeout(repairPollTimer)
  showDNSSECModal.value = false
  dnssecDomain.value = null
}

const toggleRootDNSSEC = async () => {
  const domain = dnssecDomain.value
  if (!domain) return
  if (dnssecState.value.enabled) {
    if (!confirm(`Disable DNSSEC for .${domain.domain}?\n\nRemove the DS record at the registrar FIRST and wait for its TTL to expire. Otherwise .${domain.domain} and every subdomain under it will fail validation.`)) {
      return
    }
  }
  dnssecState.value.submitting = true
  try {
    const action = dnssecState.value.enabled ? 'disable' : 'enable'
    const response = await axios.post(`/api/admin/root-domains/${domain.id}/dnssec/${action}`)
    applyRootDNSSEC(response.data)
    toast.success(response.data.message || `DNSSEC ${action}d`)
  } catch (error) {
    toast.error(error.response?.data?.error || 'DNSSEC operation failed')
  } finally {
    dnssecState.value.submitting = false
  }
}

const copyText = async (text) => {
  try {
    await navigator.clipboard.writeText(text)
    toast.success('Copied!')
  } catch {
    toast.error('Failed to copy')
  }
}

const closeModals = () => {
  showCreateModal.value = false
  showEditModal.value = false
  editingDomain.value = null
  formData.value = {
    domain: '',
    description: '',
    priority: 0,
    is_active: true,
    is_hot: false,
    is_new: false,
    is_free: true,
    min_user_level: 'normal',
    price_per_year: null,
    lifetime_price: null,
    use_default_nameservers: true,
    nameservers: [''],
  }
}
</script>
