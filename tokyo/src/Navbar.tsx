import { NavLink } from "react-router-dom";

export default function Navbar() {
  return (
    <header className="px-4 pt-8 pb-20">
      <div className="font-bold">
        <h1 className="text-center text-header">APPTRACK</h1>
        <nav className="mx-auto flex h-[36px] max-w-screen-sm items-start justify-center gap-2 text-xl">
          <NavLink
            to="/"
            className={({ isActive }) => (isActive ? "text-red-500" : "")}
          >
            <span className="flex rounded-md px-2 pb-2 transition-[padding] duration-150 ease-out hover:bg-red-500 hover:pt-2 hover:text-stone-50">
              HOME
            </span>
          </NavLink>
          <NavLink
            to="/postings"
            className={({ isActive }) => (isActive ? "text-red-500" : "")}
          >
            <span className="flex rounded-md px-2 pb-2 transition-[padding] duration-150 ease-out hover:bg-red-500 hover:pt-2 hover:text-stone-50">
              POSTINGS
            </span>
          </NavLink>
        </nav>
      </div>
    </header>
  );
}
