import { describe, it, expect } from 'vitest'
import { Endpoints } from './api'

describe('api', () => {
  describe('Endpoints', () => {
    it('should have login endpoint', () => {
      expect(Endpoints.login).toContain('/api/login')
    })

    it('should have register endpoint', () => {
      expect(Endpoints.register).toContain('/api/register')
    })

    it('should have session endpoint', () => {
      expect(Endpoints.session).toContain('/api/session')
    })

    it('should have logout endpoint', () => {
      expect(Endpoints.logout).toContain('/api/logout')
    })

    it('should have all required endpoints', () => {
      const requiredEndpoints = ['login', 'register', 'session', 'logout']
      requiredEndpoints.forEach((endpoint) => {
        expect(Endpoints).toHaveProperty(endpoint)
      })
    })
  })
})
