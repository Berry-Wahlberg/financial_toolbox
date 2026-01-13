export interface MenuItem {
  id: string
  name: string
  icon?: string
  route: string
  visible?: boolean
  order?: number
}

export interface SidebarConfig {
  items: MenuItem[]
  collapsible?: boolean
  defaultCollapsed?: boolean
}

const defaultMenuItems: MenuItem[] = [
  {
    id: 'home',
    name: 'Home',
    icon: 'House',
    route: '/',
    visible: true,
    order: 1,
  },
  {
    id: 'data-analysis',
    name: 'Data Analysis',
    icon: 'DataAnalysis',
    route: '/data-analysis',
    visible: true,
    order: 2,
  },
  {
    id: 'notes',
    name: 'Notes',
    icon: 'Document',
    route: '/notes',
    visible: true,
    order: 3,
  },
]

export const defaultSidebarConfig: SidebarConfig = {
  items: defaultMenuItems,
  collapsible: true,
  defaultCollapsed: false,
}

export const getVisibleMenuItems = (config: SidebarConfig): MenuItem[] => {
  return config.items
    .filter((item) => item.visible !== false)
    .sort((a, b) => (a.order || 0) - (b.order || 0))
}

export const getMenuItemById = (id: string, config: SidebarConfig): MenuItem | undefined => {
  return config.items.find((item) => item.id === id)
}

export const reorderMenuItems = (config: SidebarConfig, newOrder: string[]): SidebarConfig => {
  const itemMap = new Map(config.items.map((item) => [item.id, item]))
  const reorderedItems = newOrder
    .map((id) => itemMap.get(id))
    .filter((item): item is MenuItem => item !== undefined) as MenuItem[]

  return {
    ...config,
    items: reorderedItems,
  }
}

export const toggleMenuItemVisibility = (config: SidebarConfig, id: string): SidebarConfig => {
  return {
    ...config,
    items: config.items.map((item) =>
      item.id === id ? { ...item, visible: !item.visible } : item,
    ),
  }
}
