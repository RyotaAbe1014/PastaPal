import {
  IngredientCategoryCreateForm,
  IngredientCreateForm,
  IngredientListTable,
} from '@/features/ingredient'
import { Box } from '@chakra-ui/react'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_app/ingredients')({
  component: Ingredients,
})

function Ingredients() {
  return (
    <Box as="div" display={'flex'} flexDirection={'column'} gap={4}>
      <IngredientCreateForm />
      <IngredientCategoryCreateForm />
      <IngredientListTable />
    </Box>
  )
}
