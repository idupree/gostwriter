
//    Copyright 2014, Raphael Estrada
//    Author email:  <galaktor@gmx.de>
//    Project home:  <https://github.com/galaktor/gostwriter>
//    Licensed under The GPL v3 License (see README and LICENSE files)
package key

// contains key codes for use with uinput and gostwriter virtual keyboard
// 'codes.go' is intended to be generated based on 'codes.template'
// you typically should not be editing 'codes.go' but the template instead
// the constants below are directly mapped to the key codes supported by
// the linux kernel version and it's implementation of uinput
// use provided scripts to auto-generate codes for your kernel.
// [https://github.com/galaktor/gostwriter/blob/master/key/get_keycodes.sh]

/*
  #include <linux/input.h>
*/
import "C"

// a more readable and slightly more typesafe way to pass around key codes
type Code C.__u16

// all key codes that can be used with uinput and therefore gostwriter
// on the kernel that will run it. for readability, these were put into
// the 'key' package and are prefixed with 'CODE_', so they can be used
// like:
//    key.CODE_X
const (
	CODE_RESERVED         = Code(C.KEY_RESERVED)         /* 0     */
	CODE_ESC              = Code(C.KEY_ESC)              /* 1     */
	CODE_1                = Code(C.KEY_1)                /* 2     */
	CODE_2                = Code(C.KEY_2)                /* 3     */
	CODE_3                = Code(C.KEY_3)                /* 4     */
	CODE_4                = Code(C.KEY_4)                /* 5     */
	CODE_5                = Code(C.KEY_5)                /* 6     */
	CODE_6                = Code(C.KEY_6)                /* 7     */
	CODE_7                = Code(C.KEY_7)                /* 8     */
	CODE_8                = Code(C.KEY_8)                /* 9     */
	CODE_9                = Code(C.KEY_9)                /* 10    */
	CODE_0                = Code(C.KEY_0)                /* 11    */
	CODE_MINUS            = Code(C.KEY_MINUS)            /* 12    */
	CODE_EQUAL            = Code(C.KEY_EQUAL)            /* 13    */
	CODE_BACKSPACE        = Code(C.KEY_BACKSPACE)        /* 14    */
	CODE_TAB              = Code(C.KEY_TAB)              /* 15    */
	CODE_Q                = Code(C.KEY_Q)                /* 16    */
	CODE_W                = Code(C.KEY_W)                /* 17    */
	CODE_E                = Code(C.KEY_E)                /* 18    */
	CODE_R                = Code(C.KEY_R)                /* 19    */
	CODE_T                = Code(C.KEY_T)                /* 20    */
	CODE_Y                = Code(C.KEY_Y)                /* 21    */
	CODE_U                = Code(C.KEY_U)                /* 22    */
	CODE_I                = Code(C.KEY_I)                /* 23    */
	CODE_O                = Code(C.KEY_O)                /* 24    */
	CODE_P                = Code(C.KEY_P)                /* 25    */
	CODE_LEFTBRACE        = Code(C.KEY_LEFTBRACE)        /* 26    */
	CODE_RIGHTBRACE       = Code(C.KEY_RIGHTBRACE)       /* 27    */
	CODE_ENTER            = Code(C.KEY_ENTER)            /* 28    */
	CODE_LEFTCTRL         = Code(C.KEY_LEFTCTRL)         /* 29    */
	CODE_A                = Code(C.KEY_A)                /* 30    */
	CODE_S                = Code(C.KEY_S)                /* 31    */
	CODE_D                = Code(C.KEY_D)                /* 32    */
	CODE_F                = Code(C.KEY_F)                /* 33    */
	CODE_G                = Code(C.KEY_G)                /* 34    */
	CODE_H                = Code(C.KEY_H)                /* 35    */
	CODE_J                = Code(C.KEY_J)                /* 36    */
	CODE_K                = Code(C.KEY_K)                /* 37    */
	CODE_L                = Code(C.KEY_L)                /* 38    */
	CODE_SEMICOLON        = Code(C.KEY_SEMICOLON)        /* 39    */
	CODE_APOSTROPHE       = Code(C.KEY_APOSTROPHE)       /* 40    */
	CODE_GRAVE            = Code(C.KEY_GRAVE)            /* 41    */
	CODE_LEFTSHIFT        = Code(C.KEY_LEFTSHIFT)        /* 42    */
	CODE_BACKSLASH        = Code(C.KEY_BACKSLASH)        /* 43    */
	CODE_Z                = Code(C.KEY_Z)                /* 44    */
	CODE_X                = Code(C.KEY_X)                /* 45    */
	CODE_C                = Code(C.KEY_C)                /* 46    */
	CODE_V                = Code(C.KEY_V)                /* 47    */
	CODE_B                = Code(C.KEY_B)                /* 48    */
	CODE_N                = Code(C.KEY_N)                /* 49    */
	CODE_M                = Code(C.KEY_M)                /* 50    */
	CODE_COMMA            = Code(C.KEY_COMMA)            /* 51    */
	CODE_DOT              = Code(C.KEY_DOT)              /* 52    */
	CODE_SLASH            = Code(C.KEY_SLASH)            /* 53    */
	CODE_RIGHTSHIFT       = Code(C.KEY_RIGHTSHIFT)       /* 54    */
	CODE_KPASTERISK       = Code(C.KEY_KPASTERISK)       /* 55    */
	CODE_LEFTALT          = Code(C.KEY_LEFTALT)          /* 56    */
	CODE_SPACE            = Code(C.KEY_SPACE)            /* 57    */
	CODE_CAPSLOCK         = Code(C.KEY_CAPSLOCK)         /* 58    */
	CODE_F1               = Code(C.KEY_F1)               /* 59    */
	CODE_F2               = Code(C.KEY_F2)               /* 60    */
	CODE_F3               = Code(C.KEY_F3)               /* 61    */
	CODE_F4               = Code(C.KEY_F4)               /* 62    */
	CODE_F5               = Code(C.KEY_F5)               /* 63    */
	CODE_F6               = Code(C.KEY_F6)               /* 64    */
	CODE_F7               = Code(C.KEY_F7)               /* 65    */
	CODE_F8               = Code(C.KEY_F8)               /* 66    */
	CODE_F9               = Code(C.KEY_F9)               /* 67    */
	CODE_F10              = Code(C.KEY_F10)              /* 68    */
	CODE_NUMLOCK          = Code(C.KEY_NUMLOCK)          /* 69    */
	CODE_SCROLLLOCK       = Code(C.KEY_SCROLLLOCK)       /* 70    */
	CODE_KP7              = Code(C.KEY_KP7)              /* 71    */
	CODE_KP8              = Code(C.KEY_KP8)              /* 72    */
	CODE_KP9              = Code(C.KEY_KP9)              /* 73    */
	CODE_KPMINUS          = Code(C.KEY_KPMINUS)          /* 74    */
	CODE_KP4              = Code(C.KEY_KP4)              /* 75    */
	CODE_KP5              = Code(C.KEY_KP5)              /* 76    */
	CODE_KP6              = Code(C.KEY_KP6)              /* 77    */
	CODE_KPPLUS           = Code(C.KEY_KPPLUS)           /* 78    */
	CODE_KP1              = Code(C.KEY_KP1)              /* 79    */
	CODE_KP2              = Code(C.KEY_KP2)              /* 80    */
	CODE_KP3              = Code(C.KEY_KP3)              /* 81    */
	CODE_KP0              = Code(C.KEY_KP0)              /* 82    */
	CODE_KPDOT            = Code(C.KEY_KPDOT)            /* 83    */
	CODE_ZENKAKUHANKAKU   = Code(C.KEY_ZENKAKUHANKAKU)   /* 85    */
	CODE_102ND            = Code(C.KEY_102ND)            /* 86    */
	CODE_F11              = Code(C.KEY_F11)              /* 87    */
	CODE_F12              = Code(C.KEY_F12)              /* 88    */
	CODE_RO               = Code(C.KEY_RO)               /* 89    */
	CODE_KATAKANA         = Code(C.KEY_KATAKANA)         /* 90    */
	CODE_HIRAGANA         = Code(C.KEY_HIRAGANA)         /* 91    */
	CODE_HENKAN           = Code(C.KEY_HENKAN)           /* 92    */
	CODE_KATAKANAHIRAGANA = Code(C.KEY_KATAKANAHIRAGANA) /* 93    */
	CODE_MUHENKAN         = Code(C.KEY_MUHENKAN)         /* 94    */
	CODE_KPJPCOMMA        = Code(C.KEY_KPJPCOMMA)        /* 95    */
	CODE_KPENTER          = Code(C.KEY_KPENTER)          /* 96    */
	CODE_RIGHTCTRL        = Code(C.KEY_RIGHTCTRL)        /* 97    */
	CODE_KPSLASH          = Code(C.KEY_KPSLASH)          /* 98    */
	CODE_SYSRQ            = Code(C.KEY_SYSRQ)            /* 99    */
	CODE_RIGHTALT         = Code(C.KEY_RIGHTALT)         /* 100   */
	CODE_LINEFEED         = Code(C.KEY_LINEFEED)         /* 101   */
	CODE_HOME             = Code(C.KEY_HOME)             /* 102   */
	CODE_UP               = Code(C.KEY_UP)               /* 103   */
	CODE_PAGEUP           = Code(C.KEY_PAGEUP)           /* 104   */
	CODE_LEFT             = Code(C.KEY_LEFT)             /* 105   */
	CODE_RIGHT            = Code(C.KEY_RIGHT)            /* 106   */
	CODE_END              = Code(C.KEY_END)              /* 107   */
	CODE_DOWN             = Code(C.KEY_DOWN)             /* 108   */
	CODE_PAGEDOWN         = Code(C.KEY_PAGEDOWN)         /* 109   */
	CODE_INSERT           = Code(C.KEY_INSERT)           /* 110   */
	CODE_DELETE           = Code(C.KEY_DELETE)           /* 111   */
	CODE_MACRO            = Code(C.KEY_MACRO)            /* 112   */
	CODE_MUTE             = Code(C.KEY_MUTE)             /* 113   */
	CODE_VOLUMEDOWN       = Code(C.KEY_VOLUMEDOWN)       /* 114   */
	CODE_VOLUMEUP         = Code(C.KEY_VOLUMEUP)         /* 115   */
	CODE_POWER            = Code(C.KEY_POWER)            /* 116   */
	CODE_KPEQUAL          = Code(C.KEY_KPEQUAL)          /* 117   */
	CODE_KPPLUSMINUS      = Code(C.KEY_KPPLUSMINUS)      /* 118   */
	CODE_PAUSE            = Code(C.KEY_PAUSE)            /* 119   */
	CODE_SCALE            = Code(C.KEY_SCALE)            /* 120   */
	CODE_KPCOMMA          = Code(C.KEY_KPCOMMA)          /* 121   */
	CODE_HANGEUL          = Code(C.KEY_HANGEUL)          /* 122   */
	CODE_HANGUEL          = Code(C.KEY_HANGUEL)          /* KEY_HANGEUL */
	CODE_HANJA            = Code(C.KEY_HANJA)            /* 123   */
	CODE_YEN              = Code(C.KEY_YEN)              /* 124   */
	CODE_LEFTMETA         = Code(C.KEY_LEFTMETA)         /* 125   */
	CODE_RIGHTMETA        = Code(C.KEY_RIGHTMETA)        /* 126   */
	CODE_COMPOSE          = Code(C.KEY_COMPOSE)          /* 127   */
	CODE_STOP             = Code(C.KEY_STOP)             /* 128   */
	CODE_AGAIN            = Code(C.KEY_AGAIN)            /* 129   */
	CODE_PROPS            = Code(C.KEY_PROPS)            /* 130   */
	CODE_UNDO             = Code(C.KEY_UNDO)             /* 131   */
	CODE_FRONT            = Code(C.KEY_FRONT)            /* 132   */
	CODE_COPY             = Code(C.KEY_COPY)             /* 133   */
	CODE_OPEN             = Code(C.KEY_OPEN)             /* 134   */
	CODE_PASTE            = Code(C.KEY_PASTE)            /* 135   */
	CODE_FIND             = Code(C.KEY_FIND)             /* 136   */
	CODE_CUT              = Code(C.KEY_CUT)              /* 137   */
	CODE_HELP             = Code(C.KEY_HELP)             /* 138   */
	CODE_MENU             = Code(C.KEY_MENU)             /* 139   */
	CODE_CALC             = Code(C.KEY_CALC)             /* 140   */
	CODE_SETUP            = Code(C.KEY_SETUP)            /* 141   */
	CODE_SLEEP            = Code(C.KEY_SLEEP)            /* 142   */
	CODE_WAKEUP           = Code(C.KEY_WAKEUP)           /* 143   */
	CODE_FILE             = Code(C.KEY_FILE)             /* 144   */
	CODE_SENDFILE         = Code(C.KEY_SENDFILE)         /* 145   */
	CODE_DELETEFILE       = Code(C.KEY_DELETEFILE)       /* 146   */
	CODE_XFER             = Code(C.KEY_XFER)             /* 147   */
	CODE_PROG1            = Code(C.KEY_PROG1)            /* 148   */
	CODE_PROG2            = Code(C.KEY_PROG2)            /* 149   */
	CODE_WWW              = Code(C.KEY_WWW)              /* 150   */
	CODE_MSDOS            = Code(C.KEY_MSDOS)            /* 151   */
	CODE_COFFEE           = Code(C.KEY_COFFEE)           /* 152   */
	CODE_SCREENLOCK       = Code(C.KEY_SCREENLOCK)       /* KEY_COFFEE */
	CODE_DIRECTION        = Code(C.KEY_DIRECTION)        /* 153   */
	CODE_CYCLEWINDOWS     = Code(C.KEY_CYCLEWINDOWS)     /* 154   */
	CODE_MAIL             = Code(C.KEY_MAIL)             /* 155   */
	CODE_BOOKMARKS        = Code(C.KEY_BOOKMARKS)        /* 156   */
	CODE_COMPUTER         = Code(C.KEY_COMPUTER)         /* 157   */
	CODE_BACK             = Code(C.KEY_BACK)             /* 158   */
	CODE_FORWARD          = Code(C.KEY_FORWARD)          /* 159   */
	CODE_CLOSECD          = Code(C.KEY_CLOSECD)          /* 160   */
	CODE_EJECTCD          = Code(C.KEY_EJECTCD)          /* 161   */
	CODE_EJECTCLOSECD     = Code(C.KEY_EJECTCLOSECD)     /* 162   */
	CODE_NEXTSONG         = Code(C.KEY_NEXTSONG)         /* 163   */
	CODE_PLAYPAUSE        = Code(C.KEY_PLAYPAUSE)        /* 164   */
	CODE_PREVIOUSSONG     = Code(C.KEY_PREVIOUSSONG)     /* 165   */
	CODE_STOPCD           = Code(C.KEY_STOPCD)           /* 166   */
	CODE_RECORD           = Code(C.KEY_RECORD)           /* 167   */
	CODE_REWIND           = Code(C.KEY_REWIND)           /* 168   */
	CODE_PHONE            = Code(C.KEY_PHONE)            /* 169   */
	CODE_ISO              = Code(C.KEY_ISO)              /* 170   */
	CODE_CONFIG           = Code(C.KEY_CONFIG)           /* 171   */
	CODE_HOMEPAGE         = Code(C.KEY_HOMEPAGE)         /* 172   */
	CODE_REFRESH          = Code(C.KEY_REFRESH)          /* 173   */
	CODE_EXIT             = Code(C.KEY_EXIT)             /* 174   */
	CODE_MOVE             = Code(C.KEY_MOVE)             /* 175   */
	CODE_EDIT             = Code(C.KEY_EDIT)             /* 176   */
	CODE_SCROLLUP         = Code(C.KEY_SCROLLUP)         /* 177   */
	CODE_SCROLLDOWN       = Code(C.KEY_SCROLLDOWN)       /* 178   */
	CODE_KPLEFTPAREN      = Code(C.KEY_KPLEFTPAREN)      /* 179   */
	CODE_KPRIGHTPAREN     = Code(C.KEY_KPRIGHTPAREN)     /* 180   */
	CODE_NEW              = Code(C.KEY_NEW)              /* 181   */
	CODE_REDO             = Code(C.KEY_REDO)             /* 182   */
	CODE_F13              = Code(C.KEY_F13)              /* 183   */
	CODE_F14              = Code(C.KEY_F14)              /* 184   */
	CODE_F15              = Code(C.KEY_F15)              /* 185   */
	CODE_F16              = Code(C.KEY_F16)              /* 186   */
	CODE_F17              = Code(C.KEY_F17)              /* 187   */
	CODE_F18              = Code(C.KEY_F18)              /* 188   */
	CODE_F19              = Code(C.KEY_F19)              /* 189   */
	CODE_F20              = Code(C.KEY_F20)              /* 190   */
	CODE_F21              = Code(C.KEY_F21)              /* 191   */
	CODE_F22              = Code(C.KEY_F22)              /* 192   */
	CODE_F23              = Code(C.KEY_F23)              /* 193   */
	CODE_F24              = Code(C.KEY_F24)              /* 194   */
	CODE_PLAYCD           = Code(C.KEY_PLAYCD)           /* 200   */
	CODE_PAUSECD          = Code(C.KEY_PAUSECD)          /* 201   */
	CODE_PROG3            = Code(C.KEY_PROG3)            /* 202   */
	CODE_PROG4            = Code(C.KEY_PROG4)            /* 203   */
	CODE_DASHBOARD        = Code(C.KEY_DASHBOARD)        /* 204   */
	CODE_SUSPEND          = Code(C.KEY_SUSPEND)          /* 205   */
	CODE_CLOSE            = Code(C.KEY_CLOSE)            /* 206   */
	CODE_PLAY             = Code(C.KEY_PLAY)             /* 207   */
	CODE_FASTFORWARD      = Code(C.KEY_FASTFORWARD)      /* 208   */
	CODE_BASSBOOST        = Code(C.KEY_BASSBOOST)        /* 209   */
	CODE_PRINT            = Code(C.KEY_PRINT)            /* 210   */
	CODE_HP               = Code(C.KEY_HP)               /* 211   */
	CODE_CAMERA           = Code(C.KEY_CAMERA)           /* 212   */
	CODE_SOUND            = Code(C.KEY_SOUND)            /* 213   */
	CODE_QUESTION         = Code(C.KEY_QUESTION)         /* 214   */
	CODE_EMAIL            = Code(C.KEY_EMAIL)            /* 215   */
	CODE_CHAT             = Code(C.KEY_CHAT)             /* 216   */
	CODE_SEARCH           = Code(C.KEY_SEARCH)           /* 217   */
	CODE_CONNECT          = Code(C.KEY_CONNECT)          /* 218   */
	CODE_FINANCE          = Code(C.KEY_FINANCE)          /* 219   */
	CODE_SPORT            = Code(C.KEY_SPORT)            /* 220   */
	CODE_SHOP             = Code(C.KEY_SHOP)             /* 221   */
	CODE_ALTERASE         = Code(C.KEY_ALTERASE)         /* 222   */
	CODE_CANCEL           = Code(C.KEY_CANCEL)           /* 223   */
	CODE_BRIGHTNESSDOWN   = Code(C.KEY_BRIGHTNESSDOWN)   /* 224   */
	CODE_BRIGHTNESSUP     = Code(C.KEY_BRIGHTNESSUP)     /* 225   */
	CODE_MEDIA            = Code(C.KEY_MEDIA)            /* 226   */
	CODE_SWITCHVIDEOMODE  = Code(C.KEY_SWITCHVIDEOMODE)  /* 227   */
	CODE_KBDILLUMTOGGLE   = Code(C.KEY_KBDILLUMTOGGLE)   /* 228   */
	CODE_KBDILLUMDOWN     = Code(C.KEY_KBDILLUMDOWN)     /* 229   */
	CODE_KBDILLUMUP       = Code(C.KEY_KBDILLUMUP)       /* 230   */
	CODE_SEND             = Code(C.KEY_SEND)             /* 231   */
	CODE_REPLY            = Code(C.KEY_REPLY)            /* 232   */
	CODE_FORWARDMAIL      = Code(C.KEY_FORWARDMAIL)      /* 233   */
	CODE_SAVE             = Code(C.KEY_SAVE)             /* 234   */
	CODE_DOCUMENTS        = Code(C.KEY_DOCUMENTS)        /* 235   */
	CODE_BATTERY          = Code(C.KEY_BATTERY)          /* 236   */
	CODE_BLUETOOTH        = Code(C.KEY_BLUETOOTH)        /* 237   */
	CODE_WLAN             = Code(C.KEY_WLAN)             /* 238   */
	CODE_UWB              = Code(C.KEY_UWB)              /* 239   */
	CODE_UNKNOWN          = Code(C.KEY_UNKNOWN)          /* 240   */
	CODE_VIDEO_NEXT       = Code(C.KEY_VIDEO_NEXT)       /* 241   */
	CODE_VIDEO_PREV       = Code(C.KEY_VIDEO_PREV)       /* 242   */
	CODE_BRIGHTNESS_CYCLE = Code(C.KEY_BRIGHTNESS_CYCLE) /* 243   */
	CODE_BRIGHTNESS_ZERO  = Code(C.KEY_BRIGHTNESS_ZERO)  /* 244   */
	CODE_DISPLAY_OFF      = Code(C.KEY_DISPLAY_OFF)      /* 245   */
	CODE_WWAN             = Code(C.KEY_WWAN)             /* 246   */
	CODE_WIMAX            = Code(C.KEY_WIMAX)            /* KEY_WWAN */
	CODE_RFKILL           = Code(C.KEY_RFKILL)           /* 247   */
	CODE_MICMUTE          = Code(C.KEY_MICMUTE)          /* 248   */
	CODE_OK               = Code(C.KEY_OK)               /* 0x160 */
	CODE_SELECT           = Code(C.KEY_SELECT)           /* 0x161 */
	CODE_GOTO             = Code(C.KEY_GOTO)             /* 0x162 */
	CODE_CLEAR            = Code(C.KEY_CLEAR)            /* 0x163 */
	CODE_POWER2           = Code(C.KEY_POWER2)           /* 0x164 */
	CODE_OPTION           = Code(C.KEY_OPTION)           /* 0x165 */
	CODE_INFO             = Code(C.KEY_INFO)             /* 0x166 */
	CODE_TIME             = Code(C.KEY_TIME)             /* 0x167 */
	CODE_VENDOR           = Code(C.KEY_VENDOR)           /* 0x168 */
	CODE_ARCHIVE          = Code(C.KEY_ARCHIVE)          /* 0x169 */
	CODE_PROGRAM          = Code(C.KEY_PROGRAM)          /* 0x16a */
	CODE_CHANNEL          = Code(C.KEY_CHANNEL)          /* 0x16b */
	CODE_FAVORITES        = Code(C.KEY_FAVORITES)        /* 0x16c */
	CODE_EPG              = Code(C.KEY_EPG)              /* 0x16d */
	CODE_PVR              = Code(C.KEY_PVR)              /* 0x16e */
	CODE_MHP              = Code(C.KEY_MHP)              /* 0x16f */
	CODE_LANGUAGE         = Code(C.KEY_LANGUAGE)         /* 0x170 */
	CODE_TITLE            = Code(C.KEY_TITLE)            /* 0x171 */
	CODE_SUBTITLE         = Code(C.KEY_SUBTITLE)         /* 0x172 */
	CODE_ANGLE            = Code(C.KEY_ANGLE)            /* 0x173 */
	CODE_ZOOM             = Code(C.KEY_ZOOM)             /* 0x174 */
	CODE_MODE             = Code(C.KEY_MODE)             /* 0x175 */
	CODE_KEYBOARD         = Code(C.KEY_KEYBOARD)         /* 0x176 */
	CODE_SCREEN           = Code(C.KEY_SCREEN)           /* 0x177 */
	CODE_PC               = Code(C.KEY_PC)               /* 0x178 */
	CODE_TV               = Code(C.KEY_TV)               /* 0x179 */
	CODE_TV2              = Code(C.KEY_TV2)              /* 0x17a */
	CODE_VCR              = Code(C.KEY_VCR)              /* 0x17b */
	CODE_VCR2             = Code(C.KEY_VCR2)             /* 0x17c */
	CODE_SAT              = Code(C.KEY_SAT)              /* 0x17d */
	CODE_SAT2             = Code(C.KEY_SAT2)             /* 0x17e */
	CODE_CD               = Code(C.KEY_CD)               /* 0x17f */
	CODE_TAPE             = Code(C.KEY_TAPE)             /* 0x180 */
	CODE_RADIO            = Code(C.KEY_RADIO)            /* 0x181 */
	CODE_TUNER            = Code(C.KEY_TUNER)            /* 0x182 */
	CODE_PLAYER           = Code(C.KEY_PLAYER)           /* 0x183 */
	CODE_TEXT             = Code(C.KEY_TEXT)             /* 0x184 */
	CODE_DVD              = Code(C.KEY_DVD)              /* 0x185 */
	CODE_AUX              = Code(C.KEY_AUX)              /* 0x186 */
	CODE_MP3              = Code(C.KEY_MP3)              /* 0x187 */
	CODE_AUDIO            = Code(C.KEY_AUDIO)            /* 0x188 */
	CODE_VIDEO            = Code(C.KEY_VIDEO)            /* 0x189 */
	CODE_DIRECTORY        = Code(C.KEY_DIRECTORY)        /* 0x18a */
	CODE_LIST             = Code(C.KEY_LIST)             /* 0x18b */
	CODE_MEMO             = Code(C.KEY_MEMO)             /* 0x18c */
	CODE_CALENDAR         = Code(C.KEY_CALENDAR)         /* 0x18d */
	CODE_RED              = Code(C.KEY_RED)              /* 0x18e */
	CODE_GREEN            = Code(C.KEY_GREEN)            /* 0x18f */
	CODE_YELLOW           = Code(C.KEY_YELLOW)           /* 0x190 */
	CODE_BLUE             = Code(C.KEY_BLUE)             /* 0x191 */
	CODE_CHANNELUP        = Code(C.KEY_CHANNELUP)        /* 0x192 */
	CODE_CHANNELDOWN      = Code(C.KEY_CHANNELDOWN)      /* 0x193 */
	CODE_FIRST            = Code(C.KEY_FIRST)            /* 0x194 */
	CODE_LAST             = Code(C.KEY_LAST)             /* 0x195 */
	CODE_AB               = Code(C.KEY_AB)               /* 0x196 */
	CODE_NEXT             = Code(C.KEY_NEXT)             /* 0x197 */
	CODE_RESTART          = Code(C.KEY_RESTART)          /* 0x198 */
	CODE_SLOW             = Code(C.KEY_SLOW)             /* 0x199 */
	CODE_SHUFFLE          = Code(C.KEY_SHUFFLE)          /* 0x19a */
	CODE_BREAK            = Code(C.KEY_BREAK)            /* 0x19b */
	CODE_PREVIOUS         = Code(C.KEY_PREVIOUS)         /* 0x19c */
	CODE_DIGITS           = Code(C.KEY_DIGITS)           /* 0x19d */
	CODE_TEEN             = Code(C.KEY_TEEN)             /* 0x19e */
	CODE_TWEN             = Code(C.KEY_TWEN)             /* 0x19f */
	CODE_VIDEOPHONE       = Code(C.KEY_VIDEOPHONE)       /* 0x1a0 */
	CODE_GAMES            = Code(C.KEY_GAMES)            /* 0x1a1 */
	CODE_ZOOMIN           = Code(C.KEY_ZOOMIN)           /* 0x1a2 */
	CODE_ZOOMOUT          = Code(C.KEY_ZOOMOUT)          /* 0x1a3 */
	CODE_ZOOMRESET        = Code(C.KEY_ZOOMRESET)        /* 0x1a4 */
	CODE_WORDPROCESSOR    = Code(C.KEY_WORDPROCESSOR)    /* 0x1a5 */
	CODE_EDITOR           = Code(C.KEY_EDITOR)           /* 0x1a6 */
	CODE_SPREADSHEET      = Code(C.KEY_SPREADSHEET)      /* 0x1a7 */
	CODE_GRAPHICSEDITOR   = Code(C.KEY_GRAPHICSEDITOR)   /* 0x1a8 */
	CODE_PRESENTATION     = Code(C.KEY_PRESENTATION)     /* 0x1a9 */
	CODE_DATABASE         = Code(C.KEY_DATABASE)         /* 0x1aa */
	CODE_NEWS             = Code(C.KEY_NEWS)             /* 0x1ab */
	CODE_VOICEMAIL        = Code(C.KEY_VOICEMAIL)        /* 0x1ac */
	CODE_ADDRESSBOOK      = Code(C.KEY_ADDRESSBOOK)      /* 0x1ad */
	CODE_MESSENGER        = Code(C.KEY_MESSENGER)        /* 0x1ae */
	CODE_DISPLAYTOGGLE    = Code(C.KEY_DISPLAYTOGGLE)    /* 0x1af */
	CODE_SPELLCHECK       = Code(C.KEY_SPELLCHECK)       /* 0x1b0 */
	CODE_LOGOFF           = Code(C.KEY_LOGOFF)           /* 0x1b1 */
	CODE_DOLLAR           = Code(C.KEY_DOLLAR)           /* 0x1b2 */
	CODE_EURO             = Code(C.KEY_EURO)             /* 0x1b3 */
	CODE_FRAMEBACK        = Code(C.KEY_FRAMEBACK)        /* 0x1b4 */
	CODE_FRAMEFORWARD     = Code(C.KEY_FRAMEFORWARD)     /* 0x1b5 */
	CODE_CONTEXT_MENU     = Code(C.KEY_CONTEXT_MENU)     /* 0x1b6 */
	CODE_MEDIA_REPEAT     = Code(C.KEY_MEDIA_REPEAT)     /* 0x1b7 */
	CODE_10CHANNELSUP     = Code(C.KEY_10CHANNELSUP)     /* 0x1b8 */
	CODE_10CHANNELSDOWN   = Code(C.KEY_10CHANNELSDOWN)   /* 0x1b9 */
	CODE_IMAGES           = Code(C.KEY_IMAGES)           /* 0x1ba */
	CODE_DEL_EOL          = Code(C.KEY_DEL_EOL)          /* 0x1c0 */
	CODE_DEL_EOS          = Code(C.KEY_DEL_EOS)          /* 0x1c1 */
	CODE_INS_LINE         = Code(C.KEY_INS_LINE)         /* 0x1c2 */
	CODE_DEL_LINE         = Code(C.KEY_DEL_LINE)         /* 0x1c3 */
	CODE_FN               = Code(C.KEY_FN)               /* 0x1d0 */
	CODE_FN_ESC           = Code(C.KEY_FN_ESC)           /* 0x1d1 */
	CODE_FN_F1            = Code(C.KEY_FN_F1)            /* 0x1d2 */
	CODE_FN_F2            = Code(C.KEY_FN_F2)            /* 0x1d3 */
	CODE_FN_F3            = Code(C.KEY_FN_F3)            /* 0x1d4 */
	CODE_FN_F4            = Code(C.KEY_FN_F4)            /* 0x1d5 */
	CODE_FN_F5            = Code(C.KEY_FN_F5)            /* 0x1d6 */
	CODE_FN_F6            = Code(C.KEY_FN_F6)            /* 0x1d7 */
	CODE_FN_F7            = Code(C.KEY_FN_F7)            /* 0x1d8 */
	CODE_FN_F8            = Code(C.KEY_FN_F8)            /* 0x1d9 */
	CODE_FN_F9            = Code(C.KEY_FN_F9)            /* 0x1da */
	CODE_FN_F10           = Code(C.KEY_FN_F10)           /* 0x1db */
	CODE_FN_F11           = Code(C.KEY_FN_F11)           /* 0x1dc */
	CODE_FN_F12           = Code(C.KEY_FN_F12)           /* 0x1dd */
	CODE_FN_1             = Code(C.KEY_FN_1)             /* 0x1de */
	CODE_FN_2             = Code(C.KEY_FN_2)             /* 0x1df */
	CODE_FN_D             = Code(C.KEY_FN_D)             /* 0x1e0 */
	CODE_FN_E             = Code(C.KEY_FN_E)             /* 0x1e1 */
	CODE_FN_F             = Code(C.KEY_FN_F)             /* 0x1e2 */
	CODE_FN_S             = Code(C.KEY_FN_S)             /* 0x1e3 */
	CODE_FN_B             = Code(C.KEY_FN_B)             /* 0x1e4 */
	CODE_BRL_DOT1         = Code(C.KEY_BRL_DOT1)         /* 0x1f1 */
	CODE_BRL_DOT2         = Code(C.KEY_BRL_DOT2)         /* 0x1f2 */
	CODE_BRL_DOT3         = Code(C.KEY_BRL_DOT3)         /* 0x1f3 */
	CODE_BRL_DOT4         = Code(C.KEY_BRL_DOT4)         /* 0x1f4 */
	CODE_BRL_DOT5         = Code(C.KEY_BRL_DOT5)         /* 0x1f5 */
	CODE_BRL_DOT6         = Code(C.KEY_BRL_DOT6)         /* 0x1f6 */
	CODE_BRL_DOT7         = Code(C.KEY_BRL_DOT7)         /* 0x1f7 */
	CODE_BRL_DOT8         = Code(C.KEY_BRL_DOT8)         /* 0x1f8 */
	CODE_BRL_DOT9         = Code(C.KEY_BRL_DOT9)         /* 0x1f9 */
	CODE_BRL_DOT10        = Code(C.KEY_BRL_DOT10)        /* 0x1fa */
	CODE_NUMERIC_0        = Code(C.KEY_NUMERIC_0)        /* 0x200 */
	CODE_NUMERIC_1        = Code(C.KEY_NUMERIC_1)        /* 0x201 */
	CODE_NUMERIC_2        = Code(C.KEY_NUMERIC_2)        /* 0x202 */
	CODE_NUMERIC_3        = Code(C.KEY_NUMERIC_3)        /* 0x203 */
	CODE_NUMERIC_4        = Code(C.KEY_NUMERIC_4)        /* 0x204 */
	CODE_NUMERIC_5        = Code(C.KEY_NUMERIC_5)        /* 0x205 */
	CODE_NUMERIC_6        = Code(C.KEY_NUMERIC_6)        /* 0x206 */
	CODE_NUMERIC_7        = Code(C.KEY_NUMERIC_7)        /* 0x207 */
	CODE_NUMERIC_8        = Code(C.KEY_NUMERIC_8)        /* 0x208 */
	CODE_NUMERIC_9        = Code(C.KEY_NUMERIC_9)        /* 0x209 */
	CODE_NUMERIC_STAR     = Code(C.KEY_NUMERIC_STAR)     /* 0x20a */
	CODE_NUMERIC_POUND    = Code(C.KEY_NUMERIC_POUND)    /* 0x20b */
	CODE_CAMERA_FOCUS     = Code(C.KEY_CAMERA_FOCUS)     /* 0x210 */
	CODE_WPS_BUTTON       = Code(C.KEY_WPS_BUTTON)       /* 0x211 */
	CODE_TOUCHPAD_TOGGLE  = Code(C.KEY_TOUCHPAD_TOGGLE)  /* 0x212 */
	CODE_TOUCHPAD_ON      = Code(C.KEY_TOUCHPAD_ON)      /* 0x213 */
	CODE_TOUCHPAD_OFF     = Code(C.KEY_TOUCHPAD_OFF)     /* 0x214 */
	CODE_CAMERA_ZOOMIN    = Code(C.KEY_CAMERA_ZOOMIN)    /* 0x215 */
	CODE_CAMERA_ZOOMOUT   = Code(C.KEY_CAMERA_ZOOMOUT)   /* 0x216 */
	CODE_CAMERA_UP        = Code(C.KEY_CAMERA_UP)        /* 0x217 */
	CODE_CAMERA_DOWN      = Code(C.KEY_CAMERA_DOWN)      /* 0x218 */
	CODE_CAMERA_LEFT      = Code(C.KEY_CAMERA_LEFT)      /* 0x219 */
	CODE_CAMERA_RIGHT     = Code(C.KEY_CAMERA_RIGHT)     /* 0x21a */
	CODE_ATTENDANT_ON     = Code(C.KEY_ATTENDANT_ON)     /* 0x21b */
	CODE_ATTENDANT_OFF    = Code(C.KEY_ATTENDANT_OFF)    /* 0x21c */
	CODE_ATTENDANT_TOGGLE = Code(C.KEY_ATTENDANT_TOGGLE) /* 0x21d */
	CODE_LIGHTS_TOGGLE    = Code(C.KEY_LIGHTS_TOGGLE)    /* 0x21e */
	CODE_ALS_TOGGLE       = Code(C.KEY_ALS_TOGGLE)       /* 0x230 */
	CODE_MIN_INTERESTING  = Code(C.KEY_MIN_INTERESTING)  /* KEY_MUTE */
	CODE_MAX              = Code(C.KEY_MAX)              /* 0x2ff */
	CODE_CNT              = Code(C.KEY_CNT)              /* (KEY_MAX+1) */

)

// a map of all key codes for dynamic access to the
// generated codes. used to register subsets or all
// keys with uinput.
var ALL_CODES [CODE_CNT]Code = getAllCodes()

// using the linux 'input.h' key code constant 'KEY_CNT'
// which in gostwriter is represented by 'CODE_CNT'
// (note the different prefix), will iterate over all
// numerical key codes up to that limit and map them in
// 'ALL_CODES'. note that this is merely an approximation,
// and if there are gaps in the actual linux kernel key
// codes, then those gaps will not be jumped in the result
// of this function. plainly speaking, you could have key
// codes in here which don't really exist inside the kernel.
// from experience, this is unlikely to happen, so it's
// been a good heuristic so far :-)
func getAllCodes() [CODE_CNT]Code {
	result := [CODE_CNT]Code{}
	for i := 0; i < int(CODE_CNT); i++ {
		result[i] = Code(i)
	}
	return result
}
