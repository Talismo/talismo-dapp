<template>
  <v-container>
    <v-row align="center" justify="center">
      <v-col cols="6" align-self="center">
        <v-card flat class="my-3" outlined>
          <v-card-text>
            <v-container>
              <v-row align="center" justify="center">
                <v-col cols="6">
                  <h1
                    class="text-h4 center-text text-center primary--text text-uppercase"
                  >
                    Gate Search
                  </h1>
                </v-col>
              </v-row>

              <v-row align="center" justify="center">
                <v-col class="d-flex" cols="12" md="5">
                  <v-select :items="items" label="Utility Category"></v-select>
                </v-col>
                <v-col class="d-flex" cols="12" md="5">
                  <v-select :items="items" label="Sort By"></v-select>
                </v-col>
              </v-row>
              <v-row align="center" justify="center" class="mt-5">
                <v-col class="d-flex" cols="12" md="10">
                  <v-text-field
                    label="Search by Talismo username, domain or gate name"
                    :rules="rules"
                    hide-details="auto"
                  ></v-text-field>
                </v-col>
              </v-row>
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-container class="mb-5">
              <v-row align="center" justify="center">
                <v-col class="d-flex" cols="12" md="10">
                  <v-btn large block color="light-blue darken-1 white--text"
                    >Search</v-btn
                  >
                </v-col>
              </v-row>
            </v-container>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col cols="10">
        <v-container>
          <v-row>
            <v-col v-for="n in 6" :key="n" class="d-flex child-flex" cols="2">
              <v-card flat>
                <v-card-text>
                  <v-img
                    :src="`https://picsum.photos/500/300?image=${n * 5 + 10}`"
                    :lazy-src="`https://picsum.photos/10/6?image=${n * 5 + 10}`"
                    aspect-ratio="1"
                    class="grey lighten-2"
                  >
                    <template v-slot:placeholder>
                      <v-row
                        class="fill-height ma-0"
                        align="center"
                        justify="center"
                      >
                        <v-progress-circular
                          indeterminate
                          color="grey lighten-5"
                        ></v-progress-circular>
                      </v-row>
                    </template>
                  </v-img>
                </v-card-text>
                <v-card-actions>
                  <v-checkbox
                    v-model="checkbox"
                    :label="`Checkbox 1: ${checkbox.toString()}`"
                  ></v-checkbox>
                </v-card-actions>
              </v-card>
            </v-col>
          </v-row>
        </v-container>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Home",

  components: {},
  data: () => ({
    items: ["Status 1", "Status 2", "Status 3", "Status 4"],
    rules: [
      (value) => !!value || "Required.",
      (value) => (value && value.length >= 3) || "Min 3 characters",
    ],
    checkbox: true,
  }),
};
</script>
