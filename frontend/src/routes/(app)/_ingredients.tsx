import { BaseLayout } from '@/components/layout/BaseLayout'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/(app)/_ingredients')({
  component: Ingredients,
})

function Ingredients() {
  return (
    <BaseLayout path={'ingredients'}>
      <Outlet />
    </BaseLayout>
  )
}
