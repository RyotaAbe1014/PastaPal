import { BaseLayout } from '@/components/layout/BaseLayout'
import {
  IngredientCategoryCreateForm,
  IngredientCreateForm,
} from '@/features/ingredient'
import { Box } from '@chakra-ui/react'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/(app)/ingredients')({
  component: Ingredients,
})

function Ingredients() {
  return (
    <BaseLayout path={'ingredients'}>
      <Box as="div" display={'flex'} flexDirection={'column'} gap={4}>
        <IngredientCreateForm />
        <IngredientCategoryCreateForm />
        {/* TODO: 食材一覧 */}
      </Box>
    </BaseLayout>

  )
}
