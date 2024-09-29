<template>
  <div>
    <el-upload
      :action="`${getBaseUrl()}/fileUploadAndDownload/upload`"
      :before-upload="checkFile"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :show-file-list="false"
      class="upload-btn"
    >
      <el-button type="primary">上传文件</el-button>
    </el-upload>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
import { getBaseUrl } from '@/utils/format';

defineOptions({
  name: 'UploadXlsx',
});

const emit = defineEmits(['on-success']);
const path = ref(import.meta.env.VITE_BASE_API);

const fullscreenLoading = ref(false);
/**
 * 检查文件是否为XLSX格式
 */
const isXlsx = (file) => file.type === 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet';

const checkFile = (file) => {
  fullscreenLoading.value = true;
  const isLt5M = file.size / 1024 / 1024 < 5; // 5MB
  const isXlsxs = isXlsx(file);
  let pass = true;

  if (!isXlsxs) {
    ElMessage.error('请上传有效的 XLSX 文件！');
    fullscreenLoading.value = false;
    pass = false;
  }

  if (!isLt5M && isXlsx) {
    ElMessage.error('上传文件大小不能超过 5MB');
    fullscreenLoading.value = false;
    pass = false;
  }

  console.log('upload file check result: ', pass);

  return pass;
};

const uploadSuccess = (res) => {
  const { data } = res;
  if (data.file) {
    emit('on-success', data.file.url);
  }
};

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  });
  fullscreenLoading.value = false;
};
</script>