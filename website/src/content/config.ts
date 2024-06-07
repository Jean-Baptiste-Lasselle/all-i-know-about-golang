import { z, defineCollection } from 'astro:content'

// 2. Define a `type` and `schema` for each collection
const part1ChaptersCollection = defineCollection({
  type: 'content', // v2.5.0 and later
  schema: z.object({
    title: z.string(),
    summary: z.string().optional(),
    tags: z.array(z.string()),
    image: z.string().optional(),
  }),
});

// 3. Export a single `collections` object to register your collection(s)
export const collections = {
  'part1-chapters': part1ChaptersCollection,
  'part2-chapters': part1ChaptersCollection,
};
