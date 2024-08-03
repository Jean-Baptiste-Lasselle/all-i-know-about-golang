import { z, defineCollection } from 'astro:content'

// 2. Define a `type` and `schema` for each collection
const chaptersCollection = defineCollection({
  type: 'content', // v2.5.0 and later
  schema: z.object({
    title: z.string(),
    summary: z.string().optional(),
    tags: z.array(z.string()),
    image: z.string().optional(),
    index: z.number().optional().default(0),
  }),
});

// 3. Export a single `collections` object to register your collection(s)
export const collections = {
  'go-terraform-tofu': chaptersCollection,
  'questionomicon': chaptersCollection,
  'part1-chapters': chaptersCollection,
  'part2-chapters': chaptersCollection,
  'demopart-chapters': chaptersCollection,
};
