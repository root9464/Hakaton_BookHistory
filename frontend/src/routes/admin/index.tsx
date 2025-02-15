import { AdminPage } from '@/pages/AdminPage'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/admin/')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div><AdminPage /></div>
}
