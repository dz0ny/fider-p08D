interface Locale {
  text: string
  translated: number
}

const locales: { [key: string]: Locale } = {
  en: {
    text: "πΊπΈ English",
    translated: 100,
  },
  "pt-BR": {
    text: "π§π· Portuguese (Brazilian)",
    translated: 100,
  },
  "es-ES": {
    text: "πͺπΈ Spanish",
    translated: 97,
  },
  de: {
    text: "π©πͺ German",
    translated: 97,
  },
  fr: {
    text: "π«π· French",
    translated: 97,
  },
  "sv-SE": {
    text: "πΈπͺ Swedish",
    translated: 97,
  },
  nl: {
    text: "π³π± Dutch",
    translated: 97,
  },
  pl: {
    text: "π΅π± Polish",
    translated: 97,
  },
  ru: {
    text: "π·πΊ Russian",
    translated: 97,
  },
  sk: {
    text: "πΈπ° Slovak",
    translated: 97,
  },
  tr: {
    text: "πΉπ· Turkish",
    translated: 97,
  },
}

export default locales
