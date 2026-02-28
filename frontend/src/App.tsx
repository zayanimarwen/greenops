import { Routes, Route, Navigate } from 'react-router-dom'
import { useAuth } from '@/hooks/useAuth'
import { Login } from '@/pages/Login'
import { Dashboard } from '@/pages/Dashboard'
import { ClusterDetail } from '@/pages/ClusterDetail'
import { WasteExplorer } from '@/pages/WasteExplorer'
import { CarbonReport } from '@/pages/CarbonReport'
import { Savings } from '@/pages/Savings'
import { Recommendations } from '@/pages/Recommendations'
import { Simulator } from '@/pages/Simulator'
import { History } from '@/pages/History'
import { Settings } from '@/pages/Settings'
import { Tenants } from '@/pages/admin/Tenants'
import { Users } from '@/pages/admin/Users'

function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const { isAuthenticated, isLoading } = useAuth()
  if (isLoading) return <div className="flex h-screen items-center justify-center text-slate-400">Chargement...</div>
  if (!isAuthenticated) return <Navigate to="/login" replace />
  return <>{children}</>
}

export default function App() {
  return (
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route path="/*" element={
        <ProtectedRoute>
          <Routes>
            <Route path="/"                 element={<Dashboard />} />
            <Route path="/clusters/:id"     element={<ClusterDetail />} />
            <Route path="/waste"            element={<WasteExplorer />} />
            <Route path="/carbon"           element={<CarbonReport />} />
            <Route path="/savings"          element={<Savings />} />
            <Route path="/recommendations"  element={<Recommendations />} />
            <Route path="/simulator"        element={<Simulator />} />
            <Route path="/history"          element={<History />} />
            <Route path="/settings"         element={<Settings />} />
            <Route path="/admin/tenants"    element={<Tenants />} />
            <Route path="/admin/users"      element={<Users />} />
          </Routes>
        </ProtectedRoute>
      } />
    </Routes>
  )
}
