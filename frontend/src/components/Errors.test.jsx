import { describe, it, expect } from 'vitest'
import { render, screen } from '@testing-library/react'
import Errors from './Errors'

describe('Errors component', () => {
  it('should render without errors when given empty array', () => {
    render(<Errors errors={[]} />)
    const listItems = screen.queryAllByRole('listitem')
    expect(listItems).toHaveLength(0)
  })

  it('should render a single error', () => {
    render(<Errors errors={['Test error']} />)
    expect(screen.getByText('Test error')).toBeInTheDocument()
  })

  it('should render multiple errors', () => {
    const errors = ['Error 1', 'Error 2', 'Error 3']
    render(<Errors errors={errors} />)

    errors.forEach((error) => {
      expect(screen.getByText(error)).toBeInTheDocument()
    })
  })

  it('should render each error as a list item', () => {
    render(<Errors errors={['Error 1', 'Error 2']} />)
    const listItems = screen.getAllByRole('listitem')
    expect(listItems).toHaveLength(2)
  })

  it('should apply errors class to each error', () => {
    render(<Errors errors={['Test error']} />)
    const errorElement = screen.getByText('Test error')
    expect(errorElement).toHaveClass('errors')
  })
})
