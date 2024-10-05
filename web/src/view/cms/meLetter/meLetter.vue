<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip
                content="搜索范围是开始日期（包含）至结束日期（不包含）"
              >
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始日期"
            :disabled-date="
              (time) =>
                searchInfo.endCreatedAt
                  ? time.getTime() > searchInfo.endCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束日期"
            :disabled-date="
              (time) =>
                searchInfo.startCreatedAt
                  ? time.getTime() < searchInfo.startCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button
            link
            type="primary"
            icon="arrow-down"
            @click="showAllQuery = true"
            v-if="!showAllQuery"
            >展开</el-button
          >
          <el-button
            link
            type="primary"
            icon="arrow-up"
            @click="showAllQuery = false"
            v-else
            >收起</el-button
          >
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
          >删除</el-button
        >
        <ExportTemplate
          template-id="respondent"
        />
        
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>

        <el-table-column align="left" label="保单号" prop="code" width="120" />
        <el-table-column
          align="left"
          label="申请人"
          prop="applicant"
          width="120"
        />

        <el-table-column
          align="left"
          label="被申请人"
          prop="respondent"
          width="120"
        />

        <el-table-column align="left" label="法院" prop="court" width="120" />
        <el-table-column
          align="left"
          label="含税总保费"
          prop="coverage"
          width="120"
        />
        <el-table-column
          align="left"
          label="保险金额"
          prop="coverageAll"
          width="120"
        />
        <el-table-column
          align="left"
          label="开始时间"
          prop="startCreatedAt"
          width="180"
        >
          <template #default="scope">{{
            formatDate(scope.row.startCreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="结束时间"
          prop="endCreatedAt"
          width="180"
        >
          <template #default="scope">{{
            formatDate(scope.row.endCreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="保险信息"
          prop="info"
          width="120"
        />
        <el-table-column
          align="left"
          label="签署状态"
          prop="signStatus"
          width="120"
        >
          <template #default="scope">
            {{ filterDict(scope.row.signStatus, signStatusOptions) }}
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="文件下载"
          prop="templateFileUrl"
          width="120"
        />
        <el-table-column align="left" label="出单方式" prop="type" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.type, TypeOptions) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="240"
        >
          <template #default="scope">
            <el-button
              icon="download"
              type="primary"
              link
              @click="downloadFile(scope.row)"
              >下载合同</el-button
            >

            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateMeLetterFunc(scope.row)"
              >变更</el-button
            >

            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-drawer
      destroy-on-close
      size="800"
      v-model="dialogFormVisible"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === "create" ? "添加" : "修改" }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>

            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form
        :model="formData"
        label-position="top"
        ref="elFormRef"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="申请人:" prop="applicant">
          <el-select
            v-model="formData.policyholder"
            placeholder="请选择申请人"
            style="width: 100%"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in contactsOptions"
              :key="key"
              :label="item.name"
              :value="item.name"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="投保人:" prop="policyholder">
          <el-select
            v-model="formData.applicant"
            placeholder="请选择申请人"
            style="width: 100%"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in applicantOptions"
              :key="key"
              :label="item.company"
              :value="item.company"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="被申请人:" prop="respondent">
          <UploadXlsx @on-success="onFileUploaded" />
        </el-form-item>

        <el-form-item label="法院:" prop="court">
          <el-select
            v-model="formData.court"
            placeholder="请选择申请人"
            style="width: 100%"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in courtOptions"
              :key="key"
              :label="item.name"
              :value="item.name"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="含税总保费:" prop="coverage">
          <el-input
            v-model="formData.coverage"
            :clearable="true"
            placeholder="请输入含税总保费"
            :formatter="addThousandSeparatorForElementPlus"
            :parser="
              (value) =>
                removeThousandSeparatorForElementPlus(value, {
                  defaultReturn: '0',
                  decimalPlaces: 2,
                })
            "
          />
        </el-form-item>
        <el-form-item label="保险金额:" prop="coverageAll">
          <el-input
            v-model="formData.coverageAll"
            :clearable="true"
            placeholder="请输入保险金额"
            :formatter="addThousandSeparatorForElementPlus"
            :parser="
              (value) =>
                removeThousandSeparatorForElementPlus(value, {
                  defaultReturn: '0',
                  decimalPlaces: 2,
                })
            "
          />
        </el-form-item>
        <el-form-item label="开始时间:" prop="startCreatedAt">
          <el-date-picker
            v-model="formData.startCreatedAt"
            type="date"
            style="width: 30%"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="结束时间:" prop="endCreatedAt">
          <el-date-picker
            v-model="formData.endCreatedAt"
            type="date"
            style="width: 30%"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <!--            <el-form-item label="保险信息:"  prop="info" >-->
        <!--              <el-input v-model="formData.info" :clearable="true"  placeholder="请输入保险信息" />-->
        <!--            </el-form-item>-->
        <el-form-item label="出单方式:" prop="type">
          <el-select
            v-model="formData.type"
            placeholder="请选择出单方式"
            style="width: 100%"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in TypeOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createMeLetter,
  deleteMeLetter,
  deleteMeLetterByIds,
  updateMeLetter,
  findMeLetter,
  getMeLetterList,
} from "@/api/cms/meLetter";

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
  filterDataSource,
  ReturnArrImg,
  onDownloadFile,
} from "@/utils/format";
import { ElMessage, ElMessageBox,ElLoading } from "element-plus";
import { ref, reactive, watch } from "vue";
import { getContactsList } from "@/api/cms/contacts";
import { getApplicantList } from "@/api/cms/applicant";
import { getCourtList } from "@/api/cms/court";
import UploadXlsx from "@/components/upload/xlsx.vue";
import ExportTemplate from "@/components/exportExcel/exportTemplate.vue";

import {
  addThousandSeparatorForElementPlus,
  removeThousandSeparatorForElementPlus,
} from "@handsomewolf/num-utils";
import nzhcn from "nzh/cn";
defineOptions({
  name: "MeLetter",
});

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false);

// 自动化生成的字典（可能为空）以及字段
const signStatusOptions = ref([]);
const TypeOptions = ref([]);
const contactsOptions = ref({});
const applicantOptions = ref({});
const courtOptions = ref({});
const getContactsData = async () => {
  const table = await getContactsList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    contactsOptions.value = table.data.list;
  }
};
const getApplicantData = async () => {
  const table = await getApplicantList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    applicantOptions.value = table.data.list;
  }
};
const getCourtData = async () => {
  const table = await getCourtList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    courtOptions.value = table.data.list;
  }
};
const onFileUploaded = (fileName) => {
  formData.value.respondent = fileName;
  ElMessage({
    type: "success",
    message: "文件上传成功",
  });
  console.log("文件上传成功:", fileName);
};

const formData = ref({
  applicant: "",
  code: "",
  court: "",
  coverage: "",
  coveragenzh: "",
  coverageAllnzh: "",
  coverageAll: "",
  endCreatedAt: new Date(),
  info: "",
  policyholder: "",
  respondent: "",
  signStatus: "",
  startCreatedAt: new Date(),
  templateFileUrl: "",
  type: "",
});

// 验证规则
const rule = reactive({
  applicant: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  court: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  coverage: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  coverageAll: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  respondent: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "没的上传文件",
      trigger: ["input", "blur"],
    },
  ],
});

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error("请填写结束日期"));
        } else if (
          !searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt
        ) {
          callback(new Error("请填写开始日期"));
        } else if (
          searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt &&
          (searchInfo.value.startCreatedAt.getTime() ===
            searchInfo.value.endCreatedAt.getTime() ||
            searchInfo.value.startCreatedAt.getTime() >
              searchInfo.value.endCreatedAt.getTime())
        ) {
          callback(new Error("开始日期应当早于结束日期"));
        } else {
          callback();
        }
      },
      trigger: "change",
    },
  ],
});

const elFormRef = ref();
const elSearchFormRef = ref();
console.log(elFormRef.value);
// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    page.value = 1;
    pageSize.value = 10;
    getTableData();
  });
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getMeLetterList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();
getContactsData();
getApplicantData();
getCourtData();
// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
  signStatusOptions.value = await getDictFunc("signStatus");
  TypeOptions.value = await getDictFunc("Type");
};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteMeLetterFunc(row);
  });
};

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    const IDs = [];
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: "warning",
        message: "请选择要删除的数据",
      });
      return;
    }
    multipleSelection.value &&
      multipleSelection.value.map((item) => {
        IDs.push(item.ID);
      });
    const res = await deleteMeLetterByIds({ IDs });
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功",
      });
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--;
      }
      getTableData();
    }
  });
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateMeLetterFunc = async (row) => {
  const res = await findMeLetter({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteMeLetterFunc = async (row) => {
  const res = await deleteMeLetter({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    applicant: "",
    code: "",
    court: "",
    coverage: "",
    coveragenzh: "",

    coverageAll: "",
    coverageAllnzh: "",

    endCreatedAt: new Date(),
    info: "",
    policyholder: "",
    respondent: "",
    signStatus: "",
    startCreatedAt: new Date(),
    templateFileUrl: "",
    type: "",
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createMeLetter(formData.value);
        break;
      case "update":
        res = await updateMeLetter(formData.value);
        break;
      default:
        res = await createMeLetter(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      openFullScreen2();
      closeDialog();
      getTableData();
    }
  });
};

// const convertCoverageToChinese = () => {
//   formData.value.coveragenzh = nzhcn.encodeB(String(formData.value.coverage), true);
//   console.log(formData.value.coveragenzh);

// };

// // 调用这个方法来执行转换
// convertCoverageToChinese();
// 监听 coverage 的变化并转换为中文大写
watch(
  () => formData.value.coverage,
  (newCoverage) => {
    formData.value.coveragenzh = nzhcn.encodeB(newCoverage);
  }
);
watch(
  () => formData.value.coverageAll,
  (newCoverageall) => {
    formData.value.coverageAllnzh = nzhcn.encodeB(newCoverageall);
  }
);
console.log(formData.value.coveragenzh);
console.log(formData.value.coverageAllnzh);

const downloadFile = (row) => {
  onDownloadFile(row.templateFileUrl);
};
const openFullScreen2 = () => {
  const loading = ElLoading.service({
    lock: true,
    text: '签署中...',
    background: 'rgba(0, 0, 0, 0.7)',
  })
  setTimeout(() => {
    loading.close()
  }, 10000)
}
</script>

<style></style>
