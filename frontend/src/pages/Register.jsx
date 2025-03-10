import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Endpoints } from '../api'
import Errors from '../components/Errors'

const Register = () => {
  const [user, setUser] = useState({
    email: '',
    password: '',
    name: '',
  })
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [errors, setErrors] = useState([])
  const navigate = useNavigate()
  const { email, password, name } = user

  const handleChange = (e) =>
    setUser({ ...user, [e.target.name]: e.target.value })

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      setIsSubmitting(true)
      const res = await fetch(Endpoints.register, {
        method: 'POST',
        body: JSON.stringify({
          email,
          password,
          name,
        }),
        headers: {
          'Content-Type': 'application/json',
        },
      })
      const { success, errors = [] } = await res.json()

      if (success) navigate('/login')

      setErrors(errors)
    } catch (e) {
      setErrors([e.toString()])
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <div className='wrapper'>
        <h1>Register</h1>
        <input
          className='input'
          type='text'
          placeholder='Name'
          value={name}
          name='name'
          onChange={handleChange}
          required
        />
        <input
          className='input'
          type='email'
          placeholder='Email'
          value={email}
          name='email'
          onChange={handleChange}
          required
        />
        <input
          className='input'
          type='password'
          placeholder='Password'
          value={password}
          name='password'
          onChange={handleChange}
          required
        />

        <button type='submit' disabled={isSubmitting}>
          {isSubmitting ? '.....' : 'Sign Up'}
        </button>
        <br />
        <a href='/login'>{'login'}</a>
        <br />
        <Errors errors={errors} />
      </div>
    </form>
  )
}

export default Register
