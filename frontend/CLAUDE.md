# frontend

## Verify
`npx eslint src/ && npm run test:run`

## Test
```
npm run test        # watch mode
npm run test:run    # single run
```

## Structure
- Vite + React 19
- Runtime config via `public/config.js` (window.REACT_APP_API_URL)
- JSX files require `.jsx` extension
- Tests use Vitest + @testing-library/react
