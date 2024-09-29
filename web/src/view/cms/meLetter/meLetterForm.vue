<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="申请人:" prop="applicant">
          <el-input v-model="formData.applicant" :clearable="true"  placeholder="请输入申请人" />
       </el-form-item>
        <el-form-item label="保单号:" prop="code">
          <el-input v-model="formData.code" :clearable="true"  placeholder="请输入保单号" />
       </el-form-item>
        <el-form-item label="法院:" prop="court">
          <el-input v-model="formData.court" :clearable="true"  placeholder="请输入法院" />
       </el-form-item>
        <el-form-item label="含税总保费:" prop="coverage">
          <el-input v-model="formData.coverage" :clearable="true"  placeholder="请输入含税总保费" />
       </el-form-item>
        <el-form-item label="保险金额:" prop="coverageAll">
          <el-input v-model="formData.coverageAll" :clearable="true"  placeholder="请输入保险金额" />
       </el-form-item>
        <el-form-item label="结束时间:" prop="endCreatedAt">
          <el-date-picker v-model="formData.endCreatedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="文件id:" prop="fileID">
          <el-input v-model="formData.fileID" :clearable="true"  placeholder="请输入文件id" />
       </el-form-item>
        <el-form-item label="保险信息:" prop="info">
          <el-input v-model="formData.info" :clearable="true"  placeholder="请输入保险信息" />
       </el-form-item>
        <el-form-item label="投保人:" prop="policyholder">
          <el-input v-model="formData.policyholder" :clearable="true"  placeholder="请输入投保人" />
       </el-form-item>
        <el-form-item label="被申请人:" prop="respondent">
          <el-input v-model="formData.respondent" :clearable="true"  placeholder="请输入被申请人" />
       </el-form-item>
        <el-form-item label="签署状态:" prop="signStatus">
           <el-select v-model="formData.signStatus" placeholder="请选择签署状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in signStatusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="开始时间:" prop="startCreatedAt">
          <el-date-picker v-model="formData.startCreatedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="文件下载:" prop="templateFileUrl">
          <el-input v-model="formData.templateFileUrl" :clearable="true"  placeholder="请输入文件下载" />
       </el-form-item>
        <el-form-item label="出单方式:" prop="type">
           <el-select v-model="formData.type" placeholder="请选择出单方式" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in TypeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createMeLetter,
  updateMeLetter,
  findMeLetter
} from '@/api/cms/meLetter'

defineOptions({
    name: 'MeLetterForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const signStatusOptions = ref([])
const TypeOptions = ref([])
const formData = ref({
            applicant: '',
            code: '',
            court: '',
            coverage: '',
            coverageAll: '',
            endCreatedAt: new Date(),
            fileID: '',
            info: '',
            policyholder: '',
            respondent: '',
            signStatus: '',
            startCreatedAt: new Date(),
            templateFileUrl: '',
            type: '',
        })
// 验证规则
const rule = reactive({
               applicant : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               court : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               coverage : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               coverageAll : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMeLetter({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    signStatusOptions.value = await getDictFunc('signStatus')
    TypeOptions.value = await getDictFunc('Type')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createMeLetter(formData.value)
               break
             case 'update':
               res = await updateMeLetter(formData.value)
               break
             default:
               res = await createMeLetter(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
