import { describe, it, expect, beforeEach } from 'vitest'
import { createCookie, deleteCookie } from './utils'

describe('utils', () => {
  beforeEach(() => {
    // Clear cookies before each test
    document.cookie.split(';').forEach((cookie) => {
      const name = cookie.split('=')[0].trim()
      document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:01 GMT;`
    })
  })

  describe('createCookie', () => {
    it('should create a cookie with the given name and value', () => {
      createCookie('testCookie', 'testValue', 1)
      expect(document.cookie).toContain('testCookie')
      expect(document.cookie).toContain('testValue')
    })

    it('should create multiple cookies', () => {
      createCookie('cookie1', 'value1', 1)
      createCookie('cookie2', 'value2', 1)
      expect(document.cookie).toContain('cookie1')
      expect(document.cookie).toContain('cookie2')
    })
  })

  describe('deleteCookie', () => {
    it('should delete an existing cookie', () => {
      createCookie('toDelete', 'value', 1)
      expect(document.cookie).toContain('toDelete')

      deleteCookie('toDelete')
      expect(document.cookie).not.toContain('toDelete = value')
    })
  })
})
