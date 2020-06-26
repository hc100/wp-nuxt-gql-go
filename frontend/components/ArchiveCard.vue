<template>
  <div>
    <v-list-item v-for="(archive, key) in archives" :key="key" link>
      <v-list-item-action> <v-icon>mdi-page-next</v-icon> </v-list-item-action>
      <v-list-item-content>
        <v-list-item-title
          >{{ archive.year }}年{{ archive.month }}月</v-list-item-title
        >
      </v-list-item-content>
    </v-list-item>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import 'vue-apollo'
import archives from '~/apollo/queries/archives.gql'
// eslint-disable-next-line no-unused-vars
import { Archive } from '~/gql-types'

@Component({
  apollo: {
    archives: {
      prefetch: true,
      query: archives,
    },
  },
})
export default class ArchiveCard extends Vue {
  archives: Archive[] = []
}
</script>
