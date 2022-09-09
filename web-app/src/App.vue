<template>
  <v-app>
    <v-app-bar app color="light-blue darken-1" dark>
      <div class="d-flex align-center">
        <v-img
          alt="Talismo Logo"
          class="shrink mr-2"
          contain
          src="./assets/logo-talismo.png"
          transition="scale-transition"
          width="50"
        />
      </div>

      <v-spacer></v-spacer>
    </v-app-bar>

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: "App",

  data: () => ({
    //
  }),
};
</script>
