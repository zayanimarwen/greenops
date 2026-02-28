import { NavLink } from 'react-router-dom'
import { LayoutDashboard, Zap, Leaf, TrendingUp, ListChecks, PlayCircle, History, Settings, Users, Building2 } from 'lucide-react'
import { useAuthStore } from '@/store/authStore'
import { cn } from '@/lib/utils'

const nav = [
  { to: '/',                 icon: LayoutDashboard, label: 'Dashboard' },
  { to: '/waste',            icon: Zap,             label: 'Surprovisionnement' },
  { to: '/carbon',           icon: Leaf,            label: 'Impact carbone' },
  { to: '/savings',          icon: TrendingUp,      label: 'Économies' },
  { to: '/recommendations',  icon: ListChecks,      label: 'Recommandations' },
  { to: '/simulator',        icon: PlayCircle,      label: 'Simulateur' },
  { to: '/history',          icon: History,         label: 'Historique' },
  { to: '/settings',         icon: Settings,        label: 'Paramètres' },
]

const adminNav = [
  { to: '/admin/tenants', icon: Building2, label: 'Tenants' },
  { to: '/admin/users',   icon: Users,     label: 'Utilisateurs' },
]

export function Sidebar() {
  const { user } = useAuthStore()
  const isSuperAdmin = user?.roles.includes('superadmin')

  return (
    <aside className="w-64 bg-navy-900 text-white flex flex-col h-screen sticky top-0">
      <div className="p-6 border-b border-white/10">
        <div className="flex items-center gap-2">
          <span className="text-2xl">🌿</span>
          <div>
            <div className="font-bold text-white">K8s Green</div>
            <div className="text-xs text-white/50">Optimizer</div>
          </div>
        </div>
      </div>
      <nav className="flex-1 p-4 space-y-1 overflow-y-auto">
        {nav.map(({ to, icon: Icon, label }) => (
          <NavLink key={to} to={to} end={to === '/'}
            className={({ isActive }) => cn(
              'flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm transition-colors',
              isActive ? 'bg-white/10 text-white font-medium' : 'text-white/60 hover:bg-white/5 hover:text-white'
            )}>
            <Icon size={16} /> {label}
          </NavLink>
        ))}
        {isSuperAdmin && (
          <>
            <div className="pt-4 pb-1 px-3 text-xs text-white/30 uppercase tracking-wider">Admin</div>
            {adminNav.map(({ to, icon: Icon, label }) => (
              <NavLink key={to} to={to}
                className={({ isActive }) => cn(
                  'flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm transition-colors',
                  isActive ? 'bg-white/10 text-white font-medium' : 'text-white/60 hover:bg-white/5 hover:text-white'
                )}>
                <Icon size={16} /> {label}
              </NavLink>
            ))}
          </>
        )}
      </nav>
      {user && (
        <div className="p-4 border-t border-white/10">
          <div className="text-xs text-white/50 truncate">{user.email}</div>
          <div className="text-xs text-white/30">{user.roles[0]}</div>
        </div>
      )}
    </aside>
  )
}
