<template>
  <el-form
    label-position="left"
    inline
    class="proxy-table-expand"
    v-if="proxyType === 'http' || proxyType === 'https'"
  >
    <el-form-item label="Name">
      <span>{{ row.name }}</span>
    </el-form-item>
    <el-form-item label="Type">
      <span>{{ row.type }}</span>
    </el-form-item>
    <el-form-item label="Domains">
      <span>{{ row.customDomains }}</span>
    </el-form-item>
    <el-form-item label="SubDomain">
      <span>{{ row.subdomain }}</span>
    </el-form-item>
    <el-form-item label="locations">
      <span>{{ row.locations }}</span>
    </el-form-item>
    <el-form-item label="HostRewrite">
      <span>{{ row.hostHeaderRewrite }}</span>
    </el-form-item>
    <el-form-item label="Encryption">
      <span>{{ row.encryption }}</span>
    </el-form-item>
    <el-form-item label="Compression">
      <span>{{ row.compression }}</span>
    </el-form-item>
    <el-form-item label="Last Start">
      <span>{{ row.lastStartTime }}</span>
    </el-form-item>
    <el-form-item label="Last Close">
      <span>{{ row.lastCloseTime }}</span>
    </el-form-item>
  </el-form>

  <el-form label-position="left" inline class="proxy-table-expand" v-else>
    <el-form-item label="Name">
      <span>{{ row.name }}</span>
    </el-form-item>
    <el-form-item label="Type">
      <span>{{ row.type }}</span>
    </el-form-item>
    <el-form-item label="Addr">
      <span>{{ row.addr }}</span>
    </el-form-item>
    <el-form-item label="Encryption">
      <span>{{ row.encryption }}</span>
    </el-form-item>
    <el-form-item label="Compression">
      <span>{{ row.compression }}</span>
    </el-form-item>
    <el-form-item label="Last Start">
      <span>{{ row.lastStartTime }}</span>
    </el-form-item>
    <el-form-item label="Last Close">
      <span>{{ row.lastCloseTime }}</span>
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts">
defineProps<{
  row: any
  proxyType: string
}>()
</script>
