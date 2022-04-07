/*
 * This file is part of the dupman/encryptor project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>
 */

package encryptor

type Encryptor interface {
	SetPrivateKey(string) error
	PrivateKey() string
	SetPublicKey(string) error
	PublicKey() (string, error)
	GenerateKeyPair() error
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}
