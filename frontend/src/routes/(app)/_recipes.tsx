import { BaseLayout } from '@/components/layout/BaseLayout'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/(app)/_recipes')({
  component: Recipes,
})

function Recipes() {
  return (
    <BaseLayout path={'recipes'}>
      <Outlet />
    </BaseLayout>
  )
}
