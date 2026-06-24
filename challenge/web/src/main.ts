type BootState = {
  root: string
  title: string
  mode: string
  poll: number
}

export function boot(): BootState {
  return {
    root: '#app',
    title: 'atlas-node',
    mode: 'matrix',
    poll: 700
  }
}
