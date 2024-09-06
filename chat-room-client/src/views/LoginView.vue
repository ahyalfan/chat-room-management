<script setup>
import { useUsersStore } from '@/stores/users';
import Button from 'primevue/button';
import Checkbox from 'primevue/checkbox';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import InputText from 'primevue/inputtext';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const store = useUsersStore()
const router = useRouter()

const email = ref("");
const password = ref("");
const token = ref("");

// Fungsi login
const login = async (email, password) => {
  try {
    await store.loginUser(email, password);
  } catch (error) {
    console.error('Login failed:', error);
  }
};

// Fungsi untuk menangani form submit
const submitForm = async () => {
  await login(email.value, password.value);
    token.value = store.token; // Mengakses token secara langsung
    router.push({name : "room"}); // Pindah ke dashboard jika login berhasil
};
</script>

<template>
    <div>
        <!-- <input v-model="email" placeholder="email" />
        <input @keypress.enter="submitForm" v-model="password" placeholder="password" />
        <button @click="submitForm">Submit</button>
         -->

      <div v-focustrap class="w-full sm:w-80 flex flex-col gap-6">
          <IconField>
              <InputIcon>
                  <i class="pi pi-envelope" />
              </InputIcon>
              <InputText id="email" v-model="email" type="email" placeholder="Email" fluid />
          </IconField>
          <IconField>
              <InputIcon>
                  <i class="pi pi-user" />
              </InputIcon>
              <InputText id="input" v-model="password" type="password" placeholder="Password" autofocus fluid />
          </IconField>

          <div class="flex items-center gap-2">
              <Checkbox id="accept" name="accept" value="Accept" />
              <label for="accept">I agree to the terms and conditions.</label>
          </div>

          <Button type="submit" label="Submit" class="mt-2" @click="submitForm"/>
      </div>
    </div>
</template>

<style scoped>
    
</style>