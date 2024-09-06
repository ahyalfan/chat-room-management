<script setup>
import Border from '@/components/Room/Border.vue';
import { createRoom, getRooms } from '@/service/rooms';
import { useUsersStore } from '@/stores/users';
import { PrimeIcons } from '@primevue/core';
import Button from 'primevue/button';
import Column from 'primevue/column';
import DataTable from 'primevue/datatable';
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import Menu from 'primevue/menu';
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

const store = useUsersStore()
const router = useRouter();

const username = store.userName;
const token = store.token;
const visible = ref(false);

const roomId = ref("");
const newRoom = ref("");

const rooms = ref([]);

const items = ref([
    { field: 'id', header: 'Room ID' },
    { field: 'name', header: 'Room Name' },
    { field: 'users', header: 'Users' },
    { field: 'createdAt', header: 'Created At' },
]);

const saveRoom = async () => {
    try {
        const response = await createRoom(roomId.value, newRoom.value, token)
        getData(token)
    } catch (error) {
        console.log(error)
    }
    visible.value = false
}

const getData = async () => {
    try {
        const response = await getRooms(token);
        rooms.value = response.rooms;
    } catch (error) {
        console.error('Error fetching rooms:', error);
    }
};

const onSelectRow = (data) => {
    console.log('Room ID:', data.data.id);
    console.log("Room Name:", data.data.name)
    router.push(`/room/${data.data.id}/chat`);
}

onMounted(() => {
    getData();
});

</script>
<template>
    <div> 
        <div class="flex flex-row justify-center items-center gap-5">
            <h1 class="text-2xl">List of Room</h1>
            <Button label="Add Room" @click="visible = true" style="padding: 1px 6px;"/>

            <Dialog v-model:visible="visible" modal header="Add New Room" :style="{ width: '25rem' }">
                <span class="text-surface-500 dark:text-surface-400 block mb-8">Add New Room.</span>
                <div class="flex items-center gap-4 mb-4">
                    <label for="roomName" class="font-semibold w-24">Room Name</label>
                    <InputText id="roomName" v-model="newRoom" class="flex-auto" autocomplete="off" />
                </div>
                <div class="flex items-center gap-4 mb-4">
                    <label for="roomId" class="font-semibold w-24">Room Id</label>
                    <InputText id="roomId" v-model="roomId" class="flex-auto" autocomplete="off" />
                </div>

                <div class="flex justify-end gap-2">
                    <Button type="button" label="Cancel" severity="secondary" @click="visible = false"></Button>
                    <Button type="button" label="Save" @click="saveRoom"></Button>
                </div>
            </Dialog>

        </div>
        <div>
            <p>username : {{ username }}</p>
            <DataTable :value="rooms" selectionMode="single" dataKey="id" :metaKeySelection="false"
                @rowSelect="onSelectRow" tableStyle="min-width: 20rem">
                <Column field="id" header="ID" style="width: 25%"></Column>
                <Column field="name" header="Name" style="width: 25%"></Column>
            </DataTable>
        </div>
    </div>
</template>

<style scoped>
    h1 {
        color: green;
    }
</style>