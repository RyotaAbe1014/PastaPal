import { BaseLayout } from '@/components/layout/BaseLayout'
import { Outlet, createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_app')({
  component: AppLayout,
})

function AppLayout() {
  return (
    <BaseLayout>
      <Outlet />
    </BaseLayout>
  )
}
