<script setup lang="ts">
import {ref, onMounted, h} from 'vue'
import {NButton, useMessage, NTag} from "naive-ui";
import instance from "@/axios";

const message = useMessage()

let animated = ref<boolean>(false)

let searchEmail = ref<string>('')

let pageCount = ref(10)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

let dataCountOptions = [
  {
    label: '10条数据/页',
    value: 10,
  },
  {
    label: '20条数据/页',
    value: 20,
  },
  {
    label: '50条数据/页',
    value: 50,
  },
  {
    label: '100条数据/页',
    value: 100,
  },
]

interface User {
  id?: number
  role: string
  email: string
  borrowed_nums: number
  created_at: string
  updated_at: string
}

let users = ref<User[]>([])

let getAllUser = async () => {
  try {
    let {data} = await instance.get('/api/admin/v1/user', {
    params: {
      page: dataSize.value.page,
      size: dataSize.value.pageSize,
      search_email: searchEmail.value.trim(),
    }
    })
    if (data.code === 200) {
      users.value = []
      data.users.forEach((user: User) => users.value.push(user))
      pageCount.value = data.page_count
      animated.value = true
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

const columns = [
  {
    title: 'ID',
    key: 'id',
    render(row: User) {
      return h('p', {}, row.id?.toString() || '-');  // 显示ID或缺省值
    }
  },
  {
    title: '角色',
    key: 'role',
    render(row: User) {
      return h('p', {}, row.role);
    }
  },
  {
    title: '邮箱',
    key: 'email',
    render(row: User) {
      return h('p', {}, row.email);
    }
  },
  {
    title: '借阅数量',
    key: 'borrowed_nums',
    render(row: User) {
      return h(NTag, {
        type: 'success',
        bordered: false,
      }, { default: () => row.borrowed_nums.toString() });
    }
  },
  {
    title: '创建时间',
    key: 'created_at',
    render(row: User) {
      return h('p', {}, row.created_at);
    }
  },
  {
    title: '更新时间',
    key: 'updated_at',
    render(row: User) {
      return h('p', {}, row.updated_at);
    }
  }
];


onMounted(async () => {

  await getAllUser()

  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'UserMgr'
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px;">
    <n-card hoverable :embedded="true" :bordered="false">
      <div style="display: flex; flex-direction: row; justify-content: space-between; align-items: center">
        <p style="font-size: 1.1rem; font-weight: 500">用户管理</p>
<!--        <n-button size="large" type="primary" @click="editType='add'; showDrawer=true">添加图书</n-button>-->
      </div>
    </n-card>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card :embedded="true" hoverable :bordered="false" content-style="padding: 0px" style="margin-top: 20px">
        <n-data-table
            class="table"
            :columns="columns"
            :data="users"
            :pagination="false"
            :bordered="true"
            style=""
            :scroll-x="800"
        />
      </n-card>

      <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">
        <n-pagination
            size="medium"
            v-model:page.number="dataSize.page"
            :page-count="pageCount"
            @update:page="animated=false; getAllUser() "
        />
        <n-select
            style="width: 160px; margin-left: 20px"
            v-model:value.number="dataSize.pageSize"
            size="small"
            :options="dataCountOptions"
            :remote="true"
            @update:value="animated=false; dataSize.page = 1; getAllUser()"
        />
      </div>

    </div>
  </transition>

</template>

<style scoped lang="less">
.root {
 padding: 0 20px 20px 20px;
}
</style>